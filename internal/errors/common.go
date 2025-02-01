package handlederror

var (
	ErrTokenInvalid = UnauthorizedError("token is invalid").WithMessage("Session expired. Please re-Login")
	ErrUserNotFound = NotFoundError("user not found").WithMessage("User not found. Please re-Login")
)
