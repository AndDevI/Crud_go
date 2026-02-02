package users

import (
	"context"

	usersdto "crmata-go/internal/application/dto/users"
	usersservice "crmata-go/internal/application/service/users"
	domain "crmata-go/internal/domain/user"
	"crmata-go/internal/helpers"
)

func UpdateUser(ctx context.Context, service usersservice.Service, id int64, in usersdto.UpdateUserInput) (domain.User, error) {
	if id <= 0 {
		return domain.User{}, domain.ErrInvalidID
	}

	existing, err := service.Repo.FindByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	name, err := domain.NormalizeName(in.Name)
	if err != nil {
		return domain.User{}, err
	}
	email, err := domain.NormalizeEmail(in.Email)
	if err != nil {
		return domain.User{}, err
	}
	if err := domain.ValidateActive(in.Active); err != nil {
		return domain.User{}, err
	}
	if err := domain.ValidateGroupID(in.GroupID); err != nil {
		return domain.User{}, err
	}

	existing.Name = name
	existing.Email = email
	existing.Telephone = domain.NormalizeOptionalString(in.Telephone)
	existing.Image = domain.NormalizeOptionalString(in.Image)
	existing.Active = in.Active
	existing.GroupID = in.GroupID

	if in.Password != nil {
		if err := domain.ValidatePassword(*in.Password); err != nil {
			return domain.User{}, err
		}
		hashedPassword, err := helpers.HashPassword(*in.Password)
		if err != nil {
			return domain.User{}, err
		}
		existing.Password = hashedPassword
	}

	return service.Repo.Update(ctx, existing)
}
