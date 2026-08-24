package api

import (
	"net/http"

	"github.com/jb843051627/firn-signal/internal/model"
)

func (h *Handler) listScans(response http.ResponseWriter, request *http.Request) {
	values, err := h.lab.ListScans(request.Context(), request.URL.Query().Get("borehole_id"))
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, values)
}

func (h *Handler) startScan(response http.ResponseWriter, request *http.Request) {
	var input model.StartScanInput
	if err := decodeJSON(request, &input); err != nil {
		respondErr(response, err)
		return
	}
	value, err := h.lab.StartScan(request.Context(), input)
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusCreated, value)
}

func (h *Handler) getScan(response http.ResponseWriter, request *http.Request) {
	value, err := h.lab.Scan(request.Context(), request.PathValue("id"))
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, value)
}

func (h *Handler) sealScan(response http.ResponseWriter, request *http.Request) {
	value, err := h.lab.SealScan(request.Context(), request.PathValue("id"))
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, value)
}
