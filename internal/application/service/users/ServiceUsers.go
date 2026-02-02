package usersservice

import domain "crmata-go/internal/domain/user"

type Service struct {
	Repo domain.Repository
}

func NewService(repo domain.Repository) Service {
	return Service{Repo: repo}
}
