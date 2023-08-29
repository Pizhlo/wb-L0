package service

import (
	"github.com/Pizhlo/wb-L0/internal/app/storage/postgres"
)

type Service struct {
	Repo postgres.Repo
}

func New(repo postgres.Repo) *Service {
	return &Service{
		Repo: repo,
	}
}
