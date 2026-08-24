package api

import (
	"net/http"

	"github.com/jb843051627/firn-signal/internal/model"
)

func (h *Handler) recordReading(response http.ResponseWriter, request *http.Request) {
	var input model.RecordReadingInput
	if err := decodeJSON(request, &input); err != nil {
		respondErr(response, err)
		return
	}
	value, err := h.lab.RecordReading(request.Context(), input)
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusCreated, value)
}

func (h *Handler) recordReadings(response http.ResponseWriter, request *http.Request) {
	var inputs []model.RecordReadingInput
	if err := decodeJSON(request, &inputs); err != nil {
		respondErr(response, err)
		return
	}
	if err := h.lab.RecordReadings(request.Context(), inputs); err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusAccepted, map[string]int{"accepted": len(inputs)})
}
