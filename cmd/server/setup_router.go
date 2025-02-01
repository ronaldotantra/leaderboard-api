package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ronaldotantra/leaderboard-api/config"
	"github.com/ronaldotantra/leaderboard-api/internal/app"
	"github.com/ronaldotantra/leaderboard-api/internal/auth"
	handlederror "github.com/ronaldotantra/leaderboard-api/internal/errors"
	"github.com/ronaldotantra/leaderboard-api/internal/response"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func requestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New()
		c.Writer.Header().Set("X-REQUEST-ID", uuid.String())
		c.Set("X-REQUEST-ID", uuid.String())
		c.Next()
	}
}

func buildAPIRoutes(svc *app.Services, handlers *handlers) http.Handler {
	if config.IsProductionEnvironment() {
		gin.SetMode(gin.ReleaseMode)
	} else if config.IsTestEnvironment() {
		gin.SetMode(gin.TestMode)
	}

	r := gin.New()

	r.Use(corsMiddleware())
	r.Use(requestIDMiddleware())
	r.Use(gin.Logger())
	r.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
	r.Use(gin.CustomRecoveryWithWriter(os.Stderr, func(c *gin.Context, err any) {
		e := fmt.Errorf("%s", err)
		sentryHub := sentry.GetHubFromContext(c.Request.Context())
		if sentryHub != nil {
			sentryHub.CaptureException(e)
		}
		response.Failed(c, handlederror.InternalServerError(e.Error()))
	}))

	r.GET("/health-check", func(ctx *gin.Context) {
		result := svc.HealthCheckService.HealthCheck(ctx.Request.Context())
		ctx.JSON(http.StatusOK, result.MapReport())
	})

	r.POST("/api/v1/login", handlers.AuthHandlerV1.BasicLogin)
	r.POST("/api/v1/register", handlers.UserHandlerV1.Register)
	needAuthRoute := r.Group("", auth.CheckTokenMiddleware(svc.AuthService))
	needAuthRoute.GET("/api/v1/users", handlers.UserHandlerV1.GetUsers)
	needAuthRoute.POST("/api/v1/points", handlers.PointHandlerV1.InsertPoint)
	needAuthRoute.GET("/api/v1/leaderboard", handlers.PointHandlerV1.GetLeaderboard)
	return r
}
