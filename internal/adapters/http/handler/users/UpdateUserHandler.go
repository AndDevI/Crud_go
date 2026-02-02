package users

import (
	"encoding/json"
	"net/http"

	usersdto "crmata-go/internal/application/dto/users"
	usersusecases "crmata-go/internal/application/usecases/users"
	"crmata-go/internal/helpers"
)

type updateUserRequest struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  *string `json:"password"`
	Telephone *string `json:"telephone"`
	Image     *string `json:"image"`
	Active    int     `json:"active"`
	GroupID   *int64  `json:"group_id"`
}

func (h Handler) updateUser(w http.ResponseWriter, r *http.Request, id int64) {
	var req updateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, helpers.MsgInvalidJSONBody)
		return
	}

	updated, err := usersusecases.UpdateUser(r.Context(), h.service, id, usersdto.UpdateUserInput{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		Telephone: req.Telephone,
		Image:     req.Image,
		Active:    req.Active,
		GroupID:   req.GroupID,
	})
	if err != nil {
		h.handleUserError(w, err)
		return
	}

	writeSuccess(w, http.StatusOK, helpers.MsgUserUpdatedSuccess, updated)
}
