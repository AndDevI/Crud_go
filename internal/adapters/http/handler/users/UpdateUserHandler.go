package users

import (
	"encoding/json"
	"net/http"

	usersrequest "crmata-go/internal/application/request/users"
	usersusecases "crmata-go/internal/application/usecases/users"
	"crmata-go/internal/helpers"
)

func (h Handler) updateUser(w http.ResponseWriter, r *http.Request, id int64) {
	var req usersrequest.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, helpers.MsgInvalidJSONBody)
		return
	}

	updated, err := usersusecases.UpdateUser(r.Context(), h.service, id, req)
	if err != nil {
		h.handleUserError(w, err)
		return
	}

	writeSuccess(w, http.StatusOK, helpers.MsgUserUpdatedSuccess, updated)
}
