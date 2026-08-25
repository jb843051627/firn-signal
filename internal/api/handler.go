package api

import (
	"net/http"

	"github.com/jb843051627/firn-signal/internal/service"
)

type Handler struct {
	lab *service.Lab
	mux *http.ServeMux
}

func New(lab *service.Lab) http.Handler {
	h := &Handler{lab: lab, mux: http.NewServeMux()}
	h.routes()
	return h.mux
}
