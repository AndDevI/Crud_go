package users

import (
	"context"

	usersdto "crmata-go/internal/application/dto/users"
	usersservice "crmata-go/internal/application/service/users"
	domain "crmata-go/internal/domain/user"
	"crmata-go/internal/helpers"
)

func CreateUser(ctx context.Context, service usersservice.Service, in usersdto.CreateUserInput) (domain.User, error) {
	entity, err := domain.NewUser(
		in.Name,
		in.Email,
		in.Password,
		in.Telephone,
		in.Image,
		in.GroupID,
		in.Active,
	)
	if err != nil {
		return domain.User{}, err
	}

	hashedPassword, err := helpers.HashPassword(entity.Password)
	if err != nil {
		return domain.User{}, err
	}
	entity.Password = hashedPassword

	return service.Repo.Create(ctx, entity)
}
