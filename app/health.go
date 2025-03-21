package app

import (
	"net/http"

	"git.sindadsec.ir/asm/backend/clients"
	"git.sindadsec.ir/asm/backend/utils"
)

type response struct {
	Status     string `json:"status"`
	Disposable status `json:"disposable"`
}

type status struct {
	Status string `json:"status"`
	Reason string `json:"reason,omitempty"`
}

// healthCheckHandler godoc
//
//	@Summary		healthCheckHandler
//	@Description	Get server health status
//	@Tags			internal
//	@Produce		json
//	@Success		200	{object}	app.response
//	@Failure		503	{object}	app.response
//	@Router			/health [get]
func (app *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	res := &response{
		Status: "UP",
		Disposable: status{
			Status: "UP",
		},
	}
	status := http.StatusOK
	if err := clients.CheckDisposableApiHealth(app.Config.DisposableUrl); err != nil {
		res.Status = "DOWN"
		res.Disposable.Status = "DOWN"
		res.Disposable.Reason = err.Error()

		status = http.StatusServiceUnavailable
		app.Logger.Errorw(err.Error())
	}

	utils.WriteJson(w, status, res)
}
