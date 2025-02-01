package token

import (
	"context"
	"time"
)

//go:generate mockgen -source interface.go -destination ./mock/mock.go -package=tokenermock
type Tokener interface {
	GenerateToken(ctx context.Context, payload map[string]any, expiryTime *time.Duration) (string, error)
	Validate(ctx context.Context, token string) (map[string]any, error)
}
