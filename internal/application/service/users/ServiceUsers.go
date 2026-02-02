package usersservice

import userrepository "crmata-go/internal/domain/user/repository"

type Service struct {
	Repo userrepository.Repository
}

func NewService(repo userrepository.Repository) Service {
	return Service{Repo: repo}
}
