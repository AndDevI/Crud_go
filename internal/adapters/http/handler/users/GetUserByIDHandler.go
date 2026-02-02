package users

import (
	"net/http"

	usersusecases "crmata-go/internal/application/usecases/users"
	"crmata-go/internal/helpers"
)

func (h Handler) getUserByID(w http.ResponseWriter, r *http.Request, id int64) {
	user, err := usersusecases.GetUserByID(r.Context(), h.service, id)
	if err != nil {
		h.handleUserError(w, err)
		return
	}

	writeSuccess(w, http.StatusOK, helpers.MsgUserFoundSuccess, user)
}
