package user

import "context"

type Repository interface {
	SelectOneUser(ctx context.Context, input SelectOneUserInput) (output User, err error)
	SelectUsers(ctx context.Context) (output []User, err error)
	InsertUser(ctx context.Context, input RegisterPayload) error
}

//go:generate mockgen -source interface.go -destination ./mock/mock.go -package=usermock
type Service interface {
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	GetUsers(ctx context.Context) ([]User, error)
	Register(ctx context.Context, input RegisterPayload) error
}
