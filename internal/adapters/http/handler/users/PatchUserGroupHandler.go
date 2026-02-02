package users

import (
	"encoding/json"
	"errors"
	"net/http"

	usersusecases "crmata-go/internal/application/usecases/users"
	"crmata-go/internal/helpers"
)

func (h Handler) patchUserGroup(w http.ResponseWriter, r *http.Request, id int64) {
	groupID, err := decodeNullableInt64Field(r, "group_id")
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

func decodeNullableInt64Field(r *http.Request, field string) (*int64, error) {
	var raw map[string]json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&raw); err != nil {
		return nil, errors.New(helpers.MsgInvalidJSONBody)
	}

	value, ok := raw[field]
	if !ok {
		return nil, errors.New(helpers.MsgGroupIDRequired)
	}

	if string(value) == "null" {
		return nil, nil
	}

	var parsed int64
	if err := json.Unmarshal(value, &parsed); err != nil {
		return nil, errors.New(helpers.MsgGroupIDMustBeNumber)
	}

	return &parsed, nil
}
