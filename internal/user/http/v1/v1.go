package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ronaldotantra/leaderboard-api/internal/response"
	"github.com/ronaldotantra/leaderboard-api/internal/user"
)

type Handler struct {
	service user.Service
}

func New(userService user.Service) *Handler {
	return &Handler{
		service: userService,
	}
}

func (h *Handler) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := h.service.GetUsers(ctx)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success[[]user.User](c, http.StatusOK, users)
}

func (h *Handler) Register(c *gin.Context) {
	ctx := c.Request.Context()

	payload := RegisterRequest{}
	err := c.ShouldBind(&payload)
	if err != nil {
		response.Failed(c, err)
		return
	}

	err = h.service.Register(ctx, user.RegisterPayload{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	})
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success[string](c, http.StatusOK, "ok")
}
