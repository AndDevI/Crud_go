package users

import (
	"context"

	usersservice "crmata-go/internal/application/service/users"
	domain "crmata-go/internal/domain/user"
)

func GetUserByID(ctx context.Context, service usersservice.Service, id int64) (domain.User, error) {
	if id <= 0 {
		return domain.User{}, domain.ErrInvalidID
	}
	return service.Repo.FindByID(ctx, id)
}
