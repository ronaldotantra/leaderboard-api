package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
	handlederror "github.com/ronaldotantra/leaderboard-api/internal/errors"
	"github.com/ronaldotantra/leaderboard-api/internal/response"
)

func CheckTokenMiddleware(svc Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		prefix := "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Failed(c, handlederror.ErrTokenInvalid)
			c.Abort()
			return
		}

		auths := strings.Split(authHeader, prefix)
		if len(auths) < 2 {
			response.Failed(c, handlederror.ErrTokenInvalid)
			c.Abort()
			return
		}

		tokenString := auths[1]

		accessDetail, err := svc.CheckToken(c.Request.Context(), tokenString)
		if err != nil {
			response.Failed(c, err)
			c.Abort()
			return
		}

		c.Set(accessKey, accessDetail)
		c.Next()
	}
}
