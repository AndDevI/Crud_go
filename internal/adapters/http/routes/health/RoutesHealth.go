package health

import (
	"net/http"

	healthhandler "crmata-go/internal/adapters/http/handler/health"
)

func RegisterHealthRoutes(mux *http.ServeMux, handler healthhandler.Handler) {
	mux.HandleFunc("/healthz", handler.Check)
}
