package api

import "net/http"

func (h *Handler) snapshot(response http.ResponseWriter, request *http.Request) {
	value, err := h.lab.Snapshot(request.Context(), request.PathValue("id"))
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, value)
}

func (h *Handler) alerts(response http.ResponseWriter, request *http.Request) {
	value, err := h.lab.Alerts(request.Context(), request.PathValue("id"))
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, value)
}
