package httputil

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type BaseUrlKey string

const baseUrlKey BaseUrlKey = "BASE-URL"

func BaseUrlMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		baseUrl := fmt.Sprintf("%s://%s", scheme, r.Host)
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, baseUrlKey, baseUrl)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetBaseUrl(ctx context.Context) string {
	return ctx.Value(baseUrlKey).(string)
}
