package response

import (
	"github.com/gin-gonic/gin"
	handlederror "github.com/ronaldotantra/leaderboard-api/internal/errors"
)

type FailedResponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Failed(c *gin.Context, err error) {
	handledError := handlederror.Extract(err)

	c.JSON(handledError.HttpStatus, FailedResponse{
		Success: false,
		Code:    handledError.Code,
		Message: handledError.Message,
	})
}
