package api

import (
	"net/http"
	"time"
)

func (h *Handler) dailyReport(response http.ResponseWriter, request *http.Request) {
	day := time.Now().UTC()
	if raw := request.URL.Query().Get("day"); raw != "" {
		if parsed, err := time.Parse(time.RFC3339, raw); err == nil {
			day = parsed
		}
	}
	value, err := h.lab.DailyReport(request.Context(), request.URL.Query().Get("borehole_id"), day)
	if err != nil {
		respondErr(response, err)
		return
	}
	writeJSON(response, http.StatusOK, value)
}
