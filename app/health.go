package app

import (
	"net/http"

	"git.sindadsec.ir/asm/backend/utils"
)

type response struct {
	Status string `json:"status"`
}

// healthCheckHandler godoc
//
//	@Summary		healthCheckHandler
//	@Description	Get server health status
//	@Tags			internal
//	@Produce		json
//	@Success		200	{object}	app.response
//	@Router			/health [get]
func (app *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	res := &response{
		Status: "UP",
	}
	utils.WriteJson(w, http.StatusOK, res)
}
