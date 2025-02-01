package point

import "time"

type Point struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Point     int       `json:"point"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
