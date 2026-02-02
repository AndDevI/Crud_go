package users

import (
	"net/http"

	usershandler "crmata-go/internal/adapters/http/handler/users"
)

func RegisterUsersRoutes(mux *http.ServeMux, handler usershandler.Handler) {
	mux.HandleFunc("/v1/users", handler.UsersCollection)
	mux.HandleFunc("/v1/users/", handler.UsersResource)
}
