package jwttoken

import (
	"context"
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/golang-jwt/jwt"
	handlederror "github.com/ronaldotantra/leaderboard-api/internal/errors"
	"github.com/ronaldotantra/leaderboard-api/internal/logger"
	"github.com/ronaldotantra/leaderboard-api/pkg/pointer"
)

type JWTService struct {
	signMethod jwt.SigningMethod
	secret     string
}

func NewJWTService(signMethod jwt.SigningMethod, secret string) *JWTService {
	return &JWTService{signMethod, secret}
}

func (s *JWTService) GenerateToken(ctx context.Context, payload map[string]any, expiryTime *time.Duration) (string, error) {
	claims := jwt.MapClaims{}
	for key, val := range payload {
		claims[key] = val
	}

	if expiryTime == nil {
		expiryTime = pointer.Duration(time.Hour * 24)
	}
	claims["exp"] = time.Now().Add(*expiryTime).Unix()

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := jwtToken.SignedString([]byte(s.secret))
	if err != nil {
		hub := sentry.GetHubFromContext(ctx)
		if hub != nil {
			hub.CaptureException(err)
		} else {
			sentry.CaptureException(err)
		}
		return "", err
	}

	return tokenStr, nil
}

func (s *JWTService) Validate(ctx context.Context, token string) (map[string]any, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			hub := sentry.GetHubFromContext(ctx)
			if hub != nil {
				hub.CaptureException(err)
			} else {
				sentry.CaptureException(err)
			}
			return nil, err
		}
		return []byte(s.secret), nil
	})
	if err != nil {
		hub := sentry.GetHubFromContext(ctx)
		if hub != nil {
			hub.CaptureException(err)
		} else {
			sentry.CaptureException(err)
		}
		return nil, handlederror.ErrTokenInvalid
	}
	if !jwtToken.Valid {
		logger.Errorf("error token not valid: %v", err)
		return nil, handlederror.ErrTokenInvalid
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		logger.Errorf("error claims not valid: %v", err)
		return nil, handlederror.ErrTokenInvalid
	}

	return claims, nil
}
