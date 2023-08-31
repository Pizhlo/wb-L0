package service

import (
	"github.com/Pizhlo/wb-L0/internal/app/storage/cache"
	"github.com/Pizhlo/wb-L0/internal/app/storage/postgres"
)

type Service struct {
	Repo  postgres.Repo
	Cache *cache.Cache
}

func New(repo postgres.Repo, cache *cache.Cache) *Service {
	return &Service{
		Repo:  repo,
		Cache: cache,
	}
}
