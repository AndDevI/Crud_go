package users

import (
	"net/http"

	usersusecases "crmata-go/internal/application/usecases/users"
	"crmata-go/internal/helpers"
)

func (h Handler) deleteUser(w http.ResponseWriter, r *http.Request, id int64) {
	if err := usersusecases.DeleteUser(r.Context(), h.service, id); err != nil {
		h.handleUserError(w, err)
		return
	}

	writeSuccess(w, http.StatusOK, helpers.MsgUserDeletedSuccess, nil)
}
