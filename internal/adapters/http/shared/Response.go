package shared

import (
	"encoding/json"
	"net/http"

	"crmata-go/internal/helpers"
)

type APIResponse struct {
	Mensagem string `json:"mensagem"`
	Dados    any    `json:"dados,omitempty"`
}

func WriteSuccess(w http.ResponseWriter, status int, message string, data any) {
	writeJSON(w, status, APIResponse{
		Mensagem: message,
		Dados:    data,
	})
}

func WriteError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, APIResponse{Mensagem: message})
}

func WriteMethodNotAllowed(w http.ResponseWriter, allowed string) {
	w.Header().Set("Allow", allowed)
	WriteError(w, http.StatusMethodNotAllowed, helpers.MsgMethodNotAllowed)
}

func WriteNotFound(w http.ResponseWriter) {
	WriteError(w, http.StatusNotFound, helpers.MsgRouteNotFound)
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
