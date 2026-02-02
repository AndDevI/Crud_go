package users

import (
	"context"
	"strings"

	usersdto "crmata-go/internal/application/dto/users"
	usersservice "crmata-go/internal/application/service/users"
	domain "crmata-go/internal/domain/user"
)

func ListUsers(ctx context.Context, service usersservice.Service, in usersdto.ListUsersInput) ([]domain.User, error) {
	filter := domain.ListFilter{
		Search:        strings.TrimSpace(in.Search),
		SortBy:        strings.TrimSpace(in.SortBy),
		SortDirection: strings.TrimSpace(in.SortDirection),
	}
	return service.Repo.FindAll(ctx, filter)
}
