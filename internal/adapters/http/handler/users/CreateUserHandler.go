package users

import (
	"encoding/json"
	"net/http"

	usersrequest "crmata-go/internal/application/request/users"
	usersusecases "crmata-go/internal/application/usecases/users"
	"crmata-go/internal/helpers"
)

func (h Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var req usersrequest.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, helpers.MsgInvalidJSONBody)
		return
	}

	_, err := usersusecases.CreateUser(r.Context(), h.service, req)
	if err != nil {
		h.handleUserError(w, err)
		return
	}

	writeSuccess(w, http.StatusCreated, helpers.MsgUserCreatedSuccess, nil)
}
