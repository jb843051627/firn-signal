package api

import (
	"net/http"

	"github.com/jb843051627/firn-signal/internal/model"
)

func (h *Handler) recordCalibration(response http.ResponseWriter, request *http.Request) {
	var input model.RecordCalibrationInput
	if err := decodeJSON(request, &input); err != nil {
		respondErr(response, err)
		return
	}
	value, err := h.lab.RecordCalibration(request.Context(), input)
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusCreated, value)
}
