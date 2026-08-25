package api

import "net/http"

func (h *Handler) rebuildProfile(response http.ResponseWriter, request *http.Request) {
	value, err := h.lab.RebuildProfile(request.Context(), request.PathValue("id"))
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, value)
}

func (h *Handler) assessScan(response http.ResponseWriter, request *http.Request) {
	var input struct {
		Reviewer string `json:"reviewer"`
	}
	if err := decodeJSON(request, &input); err != nil {
		respondErr(response, err)
		return
	}
	value, err := h.lab.AssessScan(request.Context(), request.PathValue("id"), input.Reviewer)
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, value)
}
