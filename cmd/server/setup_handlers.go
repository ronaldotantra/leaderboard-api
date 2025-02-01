package main

import (
	"github.com/ronaldotantra/leaderboard-api/internal/app"
	authv1 "github.com/ronaldotantra/leaderboard-api/internal/auth/http/v1"
	pointv1 "github.com/ronaldotantra/leaderboard-api/internal/point/http/v1"
	userv1 "github.com/ronaldotantra/leaderboard-api/internal/user/http/v1"
)

type handlers struct {
	AuthHandlerV1  *authv1.Handler
	UserHandlerV1  *userv1.Handler
	PointHandlerV1 *pointv1.Handler
}

func setupHandlers(strg *app.Storages, services *app.Services) *handlers {
	return &handlers{
		AuthHandlerV1:  authv1.New(services.AuthService),
		UserHandlerV1:  userv1.New(services.UserService),
		PointHandlerV1: pointv1.New(services.PointService),
	}
}
