package app

import (
	"errors"
	"net/http"

	"git.sindadsec.ir/asm/backend/clients"
	"git.sindadsec.ir/asm/backend/hash"
	"git.sindadsec.ir/asm/backend/mail"
	"git.sindadsec.ir/asm/backend/models"
	"git.sindadsec.ir/asm/backend/repo"
	"git.sindadsec.ir/asm/backend/utils"
)

type registerPayload struct {
	Email        string `json:"email" validate:"required,email,non-public"`
	Password     string `json:"password" validate:"required,min=8,complex"`
	Organization string `json:"organization" validate:"required"`
}

type resendVerificationPayload struct {
	Email string `json:"email" validate:"required,email,non-public"`
}

type commonResponse struct {
	Message string `json:"message"`
}

// registerHandler godoc
//
//	@Summary		registerHandler
//	@Description	Register new user and organization
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		app.registerPayload	true	"Register payload"
//	@Success		201		{object}	app.commonResponse
//	@Failure		400		{object}	utils.errorResponse
//	@Failure		500		{object}	utils.errorResponse
//	@Router			/v1/auth/register [post]
func (app *Application) registerHandler(w http.ResponseWriter, r *http.Request) {
	var payload registerPayload
	if err := utils.ReadJson(r.Body, &payload); err != nil {
		app.Logger.Warnw(err.Error())
		utils.WriteJsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := clients.IsDisposable(app.Config.DisposableUrl, payload.Email)
	if err != nil {
		app.Logger.Errorw(err.Error())
		utils.WriteJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if *res == "true" {
		app.Logger.Warnw("disposable email detected", "email", payload.Email)
		utils.WriteJsonError(w, http.StatusBadRequest, "email address is disposable")
		return
	}

	org := &models.Organization{
		Name:  payload.Organization,
		Limit: 1,
	}
	user := &models.User{
		Email:    payload.Email,
		Password: utils.HashPassword(payload.Password),
		Enabled:  false,
		Locked:   false,
		Otp:      false,
		Role:     "ADMIN",
	}
	if err := repo.CreateUser(app.DB, app.Redis, org, user, r.Context(), app.Email); err != nil {
		switch {
		case errors.Is(err, repo.ErrDuplicateEntry):
			app.Logger.Warnw(err.Error(), "entity", "organization")
			utils.WriteJsonError(w, http.StatusBadRequest, err.Error())
		default:
			app.Logger.Errorw(err.Error())
			utils.WriteJsonError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	response := &commonResponse{
		Message: "Registration is done successfully. Please enter the 6 digit code sent to your email",
	}
	utils.WriteJson(w, http.StatusCreated, response)
}

// resendEmailVerificationHandler godoc
//
//	@Summary		resendEmailVerificationHandler
//	@Description	Resend verification email
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		app.resendVerificationPayload	true	"Resend verification payload"
//	@Success		200		{object}	app.commonResponse
//	@Failure		400		{object}	utils.errorResponse
//	@Failure		425		{object}	utils.errorResponse
//	@Failure		500		{object}	utils.errorResponse
//	@Router			/v1/auth/resend-email-verification [post]
func (app *Application) resendEmailVerificationHandler(w http.ResponseWriter, r *http.Request) {
	var payload resendVerificationPayload
	if err := utils.ReadJson(r.Body, &payload); err != nil {
		app.Logger.Warnw(err.Error())
		utils.WriteJsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := repo.GetUserByEmail(app.DB, payload.Email, r.Context())
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrRecordNotFound):
			app.Logger.Warnw(err.Error(), "email", payload.Email)
			utils.WriteJsonError(w, http.StatusBadRequest, "email address does not exists")
			return
		default:
			app.Logger.Errorw(err.Error())
			utils.WriteJsonError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if user.Enabled {
		app.Logger.Warnw("email already verified", "email", payload.Email)
		utils.WriteJsonError(w, http.StatusBadRequest, "email already verified")
		return
	}

	if err := hash.GetEmailVerificationByEmail(app.Redis, payload.Email, r.Context()); err != nil {
		switch {
		case errors.Is(err, hash.ErrNilKey):
			code, err := hash.GenerateEmailVerification(app.Redis, payload.Email, r.Context())
			if err != nil {
				app.Logger.Errorw(err.Error())
				utils.WriteJsonError(w, http.StatusInternalServerError, err.Error())
				return
			}

			if err := mail.SendRegistrationEmail(payload.Email, code, app.Email); err != nil {
				app.Logger.Errorw(err.Error())
				utils.WriteJsonError(w, http.StatusInternalServerError, err.Error())
				return
			}

			response := &commonResponse{
				Message: "email verification code has been resent",
			}
			utils.WriteJson(w, http.StatusOK, response)
			return
		default:
			app.Logger.Errorw(err.Error())
			utils.WriteJsonError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	app.Logger.Warnw("verification code not expired yet", "email", payload.Email)
	utils.WriteJsonError(w, http.StatusTooEarly, "verification code not expired yet")
}
