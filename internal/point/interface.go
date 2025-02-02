package point

import "context"

type Repository interface {
	BulkInsertPoint(ctx context.Context, input InsertPointPayload) (err error)
	GetTotalPoint(ctx context.Context, input GetTotalPointInput) (output []GetTotalPointOutput, err error)
}

//go:generate mockgen -source interface.go -destination ./mock/mock.go -package=usermock
type Service interface {
	CreatePoint(ctx context.Context, input CreatePointPayload) error
	GetLeaderboard(ctx context.Context, input GetLeaderboardPayload) (output []GetTotalPointOutput, err error)
}
