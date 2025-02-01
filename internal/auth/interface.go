package auth

import "context"

//go:generate mockgen -source interface.go -destination ./mock/mock.go -package=authmock
type Service interface {
	CheckToken(ctx context.Context, token string) (*AccessDetail, error)
	BasicLogin(ctx context.Context, payload BasicLoginPayload) (LoginResponse, error)
	Logout(ctx context.Context, token string) error
}
