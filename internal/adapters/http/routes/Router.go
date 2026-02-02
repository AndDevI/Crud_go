package routes

import (
	"net/http"

	healthhandler "crmata-go/internal/adapters/http/handler/health"
	usershandler "crmata-go/internal/adapters/http/handler/users"
	healthroutes "crmata-go/internal/adapters/http/routes/health"
	usersroutes "crmata-go/internal/adapters/http/routes/users"
)

func NewRouter(healthHandler healthhandler.Handler, usersHandler usershandler.Handler) http.Handler {
	mux := http.NewServeMux()

	healthroutes.RegisterHealthRoutes(mux, healthHandler)
	usersroutes.RegisterUsersRoutes(mux, usersHandler)

	return mux
}
