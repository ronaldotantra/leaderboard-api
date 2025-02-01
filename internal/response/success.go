package response

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponse[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

func Success[T any](c *gin.Context, httpStatus int, data T) {
	c.JSON(httpStatus, SuccessResponse[T]{
		Success: true,
		Data:    data,
	})
}
