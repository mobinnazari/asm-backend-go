package app

import (
	"errors"
	"net/http"

	"git.sindadsec.ir/asm/backend/clients"
	"git.sindadsec.ir/asm/backend/models"
	"git.sindadsec.ir/asm/backend/repo"
	"git.sindadsec.ir/asm/backend/utils"
)

type registerPayload struct {
	Email        string `json:"email" validate:"required,email,non-public"`
	Password     string `json:"password" validate:"required,min=8,complex"`
	Organization string `json:"organization" validate:"required"`
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
