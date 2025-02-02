package v1

type InsertPointRequest struct {
	UserIDs []int `json:"user_ids" binding:"required,min=2"`
}

type GetLeaderboardRequest struct {
	Month *int `form:"month" binding:"omitempty,min=1,max=12"`
	Year  *int `form:"year" binding:"omitempty"`
}
