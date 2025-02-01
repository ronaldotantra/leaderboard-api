package auth

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	handlederror "github.com/ronaldotantra/leaderboard-api/internal/errors"
	"github.com/ronaldotantra/leaderboard-api/internal/logger"
	"github.com/ronaldotantra/leaderboard-api/internal/token"
	"github.com/ronaldotantra/leaderboard-api/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	tokener     token.Tokener
	userService user.Service
}

func NewService(tokener token.Tokener, userService user.Service) Service {
	return &service{
		tokener:     tokener,
		userService: userService,
	}
}

func (s *service) CheckToken(ctx context.Context, token string) (*AccessDetail, error) {
	payload, err := s.tokener.Validate(ctx, token)
	if err != nil {
		return nil, err
	}

	userID, err := strconv.ParseInt(fmt.Sprintf("%.f", payload["user_id"]), 10, 64)
	if err != nil {
		return nil, err
	}

	return &AccessDetail{
		UserID: userID,
	}, nil
}

func (s *service) BasicLogin(ctx context.Context, payload BasicLoginPayload) (LoginResponse, error) {
	result := LoginResponse{}
	userDB, err := s.userService.GetUserByEmail(ctx, payload.Email)
	if err != nil && err != sql.ErrNoRows {
		return result, err
	}

	if err == sql.ErrNoRows {
		return result, handlederror.BadRequest("email is not registered").WithMessage("email/password is incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(payload.Password))
	if err != nil {
		return result, handlederror.BadRequest("password is incorrect").WithMessage("email/password is incorrect")
	}

	return s.login(ctx, userDB)
}

func (s *service) login(ctx context.Context, userDB user.User) (LoginResponse, error) {
	result := LoginResponse{}
	monthDuration := 30 * 24 * time.Hour

	token, err := s.tokener.GenerateToken(ctx, map[string]any{
		"user_id": userDB.ID,
		"name":    userDB.Name,
	}, &monthDuration)
	if err != nil {
		return result, err
	}
	result.Token = token
	return result, err
}

func (s *service) Logout(ctx context.Context, token string) error {
	data, err := s.tokener.Validate(ctx, token)
	if err != nil {
		return err
	}
	if data["user_id"] == nil || data["user_id"] == "" {
		logger.Errorf("error empty user_id in token: %v\n", err)
		return handlederror.ErrTokenInvalid
	}
	return nil
}
