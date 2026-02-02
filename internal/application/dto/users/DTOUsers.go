package usersdto

type CreateUserInput struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Telephone *string `json:"telephone"`
	Image     *string `json:"image"`
	Active    *int    `json:"active"`
	GroupID   *int64  `json:"group_id"`
}

type UpdateUserInput struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  *string `json:"password"`
	Telephone *string `json:"telephone"`
	Image     *string `json:"image"`
	Active    int     `json:"active"`
	GroupID   *int64  `json:"group_id"`
}

type ListUsersInput struct {
	Search        string
	SortBy        string
	SortDirection string
}
