package api

import (
	"errors"
	"net/http"

	"github.com/jb843051627/firn-signal/internal/model"
)

func statusFor(err error) int {
	if errors.Is(err, model.ErrNotFound) {
		return http.StatusNotFound
	}
	if errors.Is(err, model.ErrInvalidInput) || errors.Is(err, model.ErrInvalidState) {
		return http.StatusBadRequest
	}
	if errors.Is(err, model.ErrQualityBlock) {
		return http.StatusConflict
	}
	return http.StatusInternalServerError
}

func respondErr(response http.ResponseWriter, err error) {
	writeJSON(response, statusFor(err), map[string]string{"error": err.Error()})
}
