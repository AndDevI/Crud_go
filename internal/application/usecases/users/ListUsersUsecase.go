package users

import (
	"context"
	"strings"

	usersrequest "crmata-go/internal/application/request/users"
	usersservice "crmata-go/internal/application/service/users"
	domain "crmata-go/internal/domain/user"
	userrepository "crmata-go/internal/domain/user/repository"
)

func ListUsers(ctx context.Context, service usersservice.Service, in usersrequest.ListUsersRequest) ([]domain.User, error) {
	filter := userrepository.ListFilter{
		Search:        strings.TrimSpace(in.Search),
		SortBy:        strings.TrimSpace(in.SortBy),
		SortDirection: strings.TrimSpace(in.SortDirection),
	}
	return service.Repo.FindAll(ctx, filter)
}
