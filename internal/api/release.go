package api

import "net/http"

func (h *Handler) prepareRelease(response http.ResponseWriter, request *http.Request) {
	value, err := h.lab.PrepareRelease(request.Context(), request.PathValue("id"))
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, value)
}

func (h *Handler) publishRelease(response http.ResponseWriter, request *http.Request) {
	value, err := h.lab.PublishRelease(request.Context(), request.PathValue("id"))
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, value)
}

func (h *Handler) manifest(response http.ResponseWriter, request *http.Request) {
	value, err := h.lab.Manifest(request.Context(), request.PathValue("id"))
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, value)
}
