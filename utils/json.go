package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"git.sindadsec.ir/asm/backend/validation"
)

type errorResponse struct {
	Error string `json:"error"`
}

func WriteJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ReadJson(r io.ReadCloser, data any) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(data); err != nil {
		return err
	}
	return validation.Validate.Struct(data)
}

func WriteJsonError(w http.ResponseWriter, status int, message string) {
	data := &errorResponse{
		Error: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(*data)
}
