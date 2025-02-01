package app

import (
	"context"

	"github.com/golang-jwt/jwt"
	"github.com/ronaldotantra/leaderboard-api/config"
	"github.com/ronaldotantra/leaderboard-api/internal/auth"
	"github.com/ronaldotantra/leaderboard-api/internal/healthcheck"
	"github.com/ronaldotantra/leaderboard-api/internal/point"
	jwttoken "github.com/ronaldotantra/leaderboard-api/internal/token/jwt"
	"github.com/ronaldotantra/leaderboard-api/internal/user"
)

type Services struct {
	AuthService        auth.Service
	UserService        user.Service
	PointService       point.Service
	HealthCheckService healthcheck.Service
}

func SetupServices(ctx context.Context, strg *Storages, repositories *Repositories) *Services {
	healthCheckService := healthcheck.New(
		healthcheck.DBClientOption("leaderboard-api-db", strg.DB),
	)

	userService := user.NewService(repositories.User)
	pointService := point.NewService(repositories.Point)
	authService := auth.NewService(jwttoken.NewJWTService(jwt.SigningMethodHS256, config.AccessTokenSecretKey), userService)

	return &Services{
		HealthCheckService: healthCheckService,
		AuthService:        authService,
		UserService:        userService,
		PointService:       pointService,
	}
}
