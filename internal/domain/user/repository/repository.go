package repository

import (
	"context"
	"errors"

	"crmata-go/internal/domain/user"
	"crmata-go/internal/helpers"
)

var (
	ErrNotFound           = errors.New(helpers.MsgUserNotFound)
	ErrEmailAlreadyExists = errors.New(helpers.MsgUserEmailAlreadyExists)
	ErrGroupNotFound      = errors.New(helpers.MsgGroupNotFound)
)

type ListFilter struct {
	Search        string
	SortBy        string
	SortDirection string
}

type Repository interface {
	Create(ctx context.Context, u user.User) (user.User, error)
	FindAll(ctx context.Context, filter ListFilter) ([]user.User, error)
	FindByID(ctx context.Context, id int64) (user.User, error)
	Update(ctx context.Context, u user.User) (user.User, error)
	UpdateGroupID(ctx context.Context, id int64, groupID *int64) (user.User, error)
	Delete(ctx context.Context, id int64) error
}
