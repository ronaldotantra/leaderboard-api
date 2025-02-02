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

func (s *service) GetLeaderboard(ctx context.Context, input GetLeaderboardPayload) (output []GetTotalPointOutput, err error) {
	now := time.Now()
	month := now.Month()
	year := now.Year()
	if input.Month != nil {
		month = time.Month(*input.Month)
	}
	if input.Year != nil {
		year = *input.Year
	}
	jakarta := time.FixedZone("Asia/Jakarta", 7*60*60)
	return s.repo.GetTotalPoint(ctx, GetTotalPointInput{
		StartDate: time.Date(year, month, 1, 0, 0, 0, 0, jakarta),
		EndDate:   time.Date(year, month+1, 1, 0, 0, 0, 0, jakarta),
	})
}
