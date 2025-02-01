package v1

type InsertPointRequest struct {
	UserIDs []int `json:"user_ids" binding:"required,min=2"`
}

type GetLeaderboardRequest struct {
	Month *int `json:"month" binding:"required,min=1,max=12"`
	Year  *int `json:"year" binding:"required`
}
