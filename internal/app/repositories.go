package app

import (
	"github.com/ronaldotantra/go-atomic"
	"github.com/ronaldotantra/leaderboard-api/internal/point"
	pointrepository "github.com/ronaldotantra/leaderboard-api/internal/point/repository"
	"github.com/ronaldotantra/leaderboard-api/internal/user"
	userrepository "github.com/ronaldotantra/leaderboard-api/internal/user/repository"
)

type Repositories struct {
	User  user.Repository
	Point point.Repository
}

func SetupRepositories(strg *Storages) *Repositories {
	atomicExecutor := atomic.New(strg.DB)
	return &Repositories{
		User:  userrepository.New(atomicExecutor),
		Point: pointrepository.New(atomicExecutor),
	}
}
