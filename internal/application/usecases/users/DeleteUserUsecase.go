package users

import (
	"context"

	usersservice "crmata-go/internal/application/service/users"
	domain "crmata-go/internal/domain/user"
)

func DeleteUser(ctx context.Context, service usersservice.Service, id int64) error {
	if id <= 0 {
		return domain.ErrInvalidID
	}
	return service.Repo.Delete(ctx, id)
}
