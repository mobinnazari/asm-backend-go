package app

import (
	"net/http"

	"git.sindadsec.ir/asm/backend/utils"
)

type response struct {
	Status string `json:"status"`
}

func (app *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	res := &response{
		Status: "UP",
	}
	utils.WriteJson(w, http.StatusOK, res)
}
