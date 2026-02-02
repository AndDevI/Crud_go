package users

import (
	"net/http"

	usersdto "crmata-go/internal/application/dto/users"
	usersusecases "crmata-go/internal/application/usecases/users"
	"crmata-go/internal/helpers"
)

func (h Handler) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := usersusecases.ListUsers(r.Context(), h.service, usersdto.ListUsersInput{
		Search:        r.URL.Query().Get("search"),
		SortBy:        r.URL.Query().Get("sort_by"),
		SortDirection: r.URL.Query().Get("sort_direction"),
	})
	if err != nil {
		h.handleUserError(w, err)
		return
	}

	writeSuccess(w, http.StatusOK, helpers.MsgUsersListedSuccess, users)
}
