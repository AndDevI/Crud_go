package usersrequest

import "encoding/json"

type CreateUserRequest struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Telephone *string `json:"telephone"`
	Image     *string `json:"image"`
	Active    *int    `json:"active"`
	GroupID   *int64  `json:"group_id"`
}

type UpdateUserRequest struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  *string `json:"password"`
	Telephone *string `json:"telephone"`
	Image     *string `json:"image"`
	Active    int     `json:"active"`
	GroupID   *int64  `json:"group_id"`
}

type ListUsersRequest struct {
	Search        string
	SortBy        string
	SortDirection string
}

type UpdateUserGroupRequest struct {
	GroupID *json.RawMessage `json:"group_id"`
}
