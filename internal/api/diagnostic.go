package api

import "net/http"

func (h *Handler) diagnostics(response http.ResponseWriter, request *http.Request) {
	value, err := h.lab.Diagnostics(request.Context(), request.PathValue("id"))
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, value)
}
