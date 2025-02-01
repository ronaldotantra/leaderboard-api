package auth

type AccessDetail struct {
	UserID int64
}

type LoginResponse struct {
	Token string
}
