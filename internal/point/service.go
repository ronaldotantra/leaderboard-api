package point

import (
	"context"
	"time"

	handlederror "github.com/ronaldotantra/leaderboard-api/internal/errors"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreatePoint(ctx context.Context, input CreatePointPayload) error {
	items := []InsertPointItem{}
	uniqueUserId := make(map[int]bool)
	for _, userID := range input.UserIDs {
		if uniqueUserId[userID] {
			continue
		}
		uniqueUserId[userID] = true
	}
	lenUniqueUserIds := len(input.UserIDs)
	if lenUniqueUserIds != len(uniqueUserId) {
		return handlederror.BadRequest("user_id_duplicate").WithMessage("user_id_duplicate")
	}
	currentScore := len(input.UserIDs) / 2
	isEven := len(input.UserIDs)%2 == 0
	now := time.Now()
	for _, userID := range input.UserIDs {
		if isEven && currentScore == 0 {
			currentScore--
		}
		items = append(items, InsertPointItem{
			UserID: userID,
			Point:  currentScore,
			Date:   now,
		})
		currentScore--
	}
	return s.repo.BulkInsertPoint(ctx, InsertPointPayload{
		Items:   items,
		UserIDs: input.UserIDs,
	})
}

func (s *service) GetLeaderboard(ctx context.Context, input GetLeaderboardPayload) (err error) {
	return
}
