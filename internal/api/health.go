package api

import "net/http"

func (h *Handler) health(response http.ResponseWriter, request *http.Request) {
	value, err := h.lab.Health(request.Context())
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, value)
}
