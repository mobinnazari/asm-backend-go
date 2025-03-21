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
		utils.WriteJsonError(w, http.StatusServiceUnavailable, err.Error())
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
	if err := repo.CreateUser(app.DB, org, r.Context()); err != nil {
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

	utils.WriteJson(w, http.StatusNoContent, nil)
}
