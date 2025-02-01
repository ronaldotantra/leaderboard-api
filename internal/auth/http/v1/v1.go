package authv1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ronaldotantra/leaderboard-api/internal/auth"
	"github.com/ronaldotantra/leaderboard-api/internal/response"
)

type Handler struct {
	service auth.Service
}

func New(authService auth.Service) *Handler {
	return &Handler{
		service: authService,
	}
}

func (h *Handler) BasicLogin(c *gin.Context) {
	ctx := c.Request.Context()

	payload := BasicLoginPayload{}
	err := c.ShouldBind(&payload)
	if err != nil {
		response.Failed(c, err)
		return
	}

	result, err := h.service.BasicLogin(ctx, auth.BasicLoginPayload{
		Email:    strings.ToLower(payload.Email),
		Password: payload.Password,
	})
	if err != nil {
		response.Failed(c, err)
		return
	}
	resp := loginResponse{
		Token: result.Token,
	}

	response.Success[loginResponse](c, http.StatusOK, resp)
}

func (h *Handler) Logout(c *gin.Context) {
	ctx := c.Request.Context()

	token := getBearer(c.Request)
	err := h.service.Logout(ctx, token)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success[string](c, http.StatusOK, "ok")
}
