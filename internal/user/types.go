package user

type SelectOneUserInput struct {
	UserID *int64
	Email  *string
}

type RegisterPayload struct {
	Name     string
	Email    string
	Password string
}
