package user

import (
	"context"
	"errors"
)

var (
	ErrNotFound           = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("user email already exists")
	ErrGroupNotFound      = errors.New("group not found")
)

type ListFilter struct {
	Search        string
	SortBy        string
	SortDirection string
}

type Repository interface {
	Create(ctx context.Context, u User) (User, error)
	FindAll(ctx context.Context, filter ListFilter) ([]User, error)
	FindByID(ctx context.Context, id int64) (User, error)
	Update(ctx context.Context, u User) (User, error)
	UpdateGroupID(ctx context.Context, id int64, groupID *int64) (User, error)
	Delete(ctx context.Context, id int64) error
}
