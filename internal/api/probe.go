package api

import (
	"net/http"

	"github.com/jb843051627/firn-signal/internal/model"
)

func (h *Handler) listProbes(response http.ResponseWriter, request *http.Request) {
	values, err := h.lab.ListProbes(request.Context(), request.URL.Query().Get("borehole_id"))
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, values)
}

func (h *Handler) registerProbe(response http.ResponseWriter, request *http.Request) {
	var input model.RegisterProbeInput
	if err := decodeJSON(request, &input); err != nil {
		respondErr(response, err)
		return
	}
	value, err := h.lab.RegisterProbe(request.Context(), input)
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusCreated, value)
}

func (h *Handler) removeProbe(response http.ResponseWriter, request *http.Request) {
	value, err := h.lab.RemoveProbe(request.Context(), request.PathValue("id"))
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, value)
}
