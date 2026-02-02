package health

import (
	"net/http"

	"crmata-go/internal/adapters/http/shared"
	"crmata-go/internal/helpers"
)

type Handler struct{}

func New() Handler {
	return Handler{}
}

func (h Handler) Check(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		shared.WriteMethodNotAllowed(w, "GET")
		return
	}

	shared.WriteSuccess(w, http.StatusOK, helpers.MsgHealthOK, map[string]string{"status": "ok"})
}
