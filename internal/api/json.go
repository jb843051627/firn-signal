package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jb843051627/firn-signal/internal/model"
)

func decodeJSON(request *http.Request, target any) error {
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(target)
}

func writeJSON(response http.ResponseWriter, status int, value any) {
	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	response.WriteHeader(status)
	_ = json.NewEncoder(response).Encode(value)
}

func writeError(response http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	if errors.Is(err, model.ErrNotFound) {
		status = http.StatusNotFound
	}
	if errors.Is(err, model.ErrInvalidInput) || errors.Is(err, model.ErrInvalidState) {
		status = http.StatusBadRequest
	}
	writeJSON(response, status, map[string]string{"error": err.Error()})
}
