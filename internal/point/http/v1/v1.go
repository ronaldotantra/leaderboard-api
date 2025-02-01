package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	handlederror "github.com/ronaldotantra/leaderboard-api/internal/errors"
	"github.com/ronaldotantra/leaderboard-api/internal/point"
	"github.com/ronaldotantra/leaderboard-api/internal/response"
)

type Handler struct {
	service point.Service
}

func New(pointService point.Service) *Handler {
	return &Handler{
		service: pointService,
	}
}

func (h *Handler) InsertPoint(c *gin.Context) {
	ctx := c.Request.Context()

	payload := InsertPointRequest{}
	err := c.ShouldBind(&payload)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			response.Failed(c, handlederror.BadRequest(validationErrors.Error()))
		} else {
			response.Failed(c, err)
		}
		return
	}

	err = h.service.CreatePoint(ctx, point.CreatePointPayload{
		UserIDs: payload.UserIDs,
	})
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success[string](c, http.StatusOK, "ok")
}

func (h *Handler) GetLeaderboard(c *gin.Context) {
	ctx := c.Request.Context()

	payload := GetLeaderboardRequest{}
	err := c.ShouldBind(&payload)
	if err != nil {
		response.Failed(c, err)
		return
	}

	err = h.service.GetLeaderboard(ctx, point.GetLeaderboardPayload{
		Month: payload.Month,
		Year:  payload.Year,
	})
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success[string](c, http.StatusOK, "ok")
}
