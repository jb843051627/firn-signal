package api

import (
	"net/http"

	"github.com/jb843051627/firn-signal/internal/model"
)

func (h *Handler) listBoreholes(response http.ResponseWriter, request *http.Request) {
	values, err := h.lab.ListBoreholes(request.Context())
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, values)
}

func (h *Handler) createBorehole(response http.ResponseWriter, request *http.Request) {
	var input model.CreateBoreholeInput
	_ = decodeJSON(request, &input)
	value, err := h.lab.CreateBorehole(request.Context(), input)
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusCreated, value)
}

func (h *Handler) getBorehole(response http.ResponseWriter, request *http.Request) {
	value, err := h.lab.Borehole(request.Context(), request.PathValue("id"))
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, value)
}

func (h *Handler) archiveBorehole(response http.ResponseWriter, request *http.Request) {
	value, err := h.lab.ArchiveBorehole(request.Context(), request.PathValue("id"))
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, value)
}
