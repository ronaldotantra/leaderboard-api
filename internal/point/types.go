package point

import "time"

type CreatePointPayload struct {
	UserIDs []int
}

type InsertPointPayload struct {
	Items   []InsertPointItem
	UserIDs []int
}

type InsertPointItem struct {
	UserID int
	Point  int
	Date   time.Time
}

type GetLeaderboardPayload struct {
	Month *int
	Year  *int
}

type GetLeaderboardInput struct {
	Month int
	Year  int
}

type GetTotalPointInput struct {
	StartDate time.Time
	EndDate   time.Time
}

type GetTotalPointOutput struct {
	Name       string
	TotalPoint int
}
