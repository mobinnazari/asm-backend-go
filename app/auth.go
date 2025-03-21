package app

import (
	"net/http"

	"git.sindadsec.ir/asm/backend/clients"
	"git.sindadsec.ir/asm/backend/utils"
)

type registerPayload struct {
	Email string `json:"email" validate:"required,email,non-public"`
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

	utils.WriteJson(w, http.StatusNoContent, nil)
}
