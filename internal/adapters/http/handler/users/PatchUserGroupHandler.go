package users

import (
	"encoding/json"
	"errors"
	"net/http"

	usersrequest "crmata-go/internal/application/request/users"
	usersusecases "crmata-go/internal/application/usecases/users"
	"crmata-go/internal/helpers"
)

func (h Handler) patchUserGroup(w http.ResponseWriter, r *http.Request, id int64) {
	groupID, err := decodeNullableGroupID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	updated, err := usersusecases.UpdateUserGroupID(r.Context(), h.service, id, groupID)
	if err != nil {
		h.handleUserError(w, err)
		return
	}

	writeSuccess(w, http.StatusOK, helpers.MsgUserGroupUpdated, updated)
}

func decodeNullableGroupID(r *http.Request) (*int64, error) {
	var req usersrequest.UpdateUserGroupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.New(helpers.MsgInvalidJSONBody)
	}

	if req.GroupID == nil {
		return nil, errors.New(helpers.MsgGroupIDRequired)
	}

	if string(*req.GroupID) == "null" {
		return nil, nil
	}

	var parsed int64
	if err := json.Unmarshal(*req.GroupID, &parsed); err != nil {
		return nil, errors.New(helpers.MsgGroupIDMustBeNumber)
	}

	return &parsed, nil
}
