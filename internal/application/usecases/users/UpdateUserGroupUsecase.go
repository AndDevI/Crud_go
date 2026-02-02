package users

import (
	"context"

	usersservice "crmata-go/internal/application/service/users"
	domain "crmata-go/internal/domain/user"
)

func UpdateUserGroupID(ctx context.Context, service usersservice.Service, id int64, groupID *int64) (domain.User, error) {
	if id <= 0 {
		return domain.User{}, domain.ErrInvalidID
	}
	if err := domain.ValidateGroupID(groupID); err != nil {
		return domain.User{}, err
	}
	return service.Repo.UpdateGroupID(ctx, id, groupID)
}
