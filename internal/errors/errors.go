package handlederror

import "net/http"

const (
	DefaultErrorMessage     = "Currently the server can not process your request. Please try again later."
	BadRequestCode          = "BAD_REQUEST"
	InternalServerErrorCode = "INTERNAL_SERVER_ERROR"
	ValidationErrorCode     = "VALIDATION_ERROR"
	UnauthorizedErrorCode   = "UNAUTHORIZED"
	NotFoundErrorCode       = "NOT_FOUND"
)

func InternalServerError(detail string) HandledError {
	return HandledError{
		HttpStatus: http.StatusInternalServerError,
		Code:       InternalServerErrorCode,
		Message:    DefaultErrorMessage,
		Detail:     detail,
	}
}

func ValidationError(detail string) HandledError {
	return HandledError{
		HttpStatus: http.StatusBadRequest,
		Code:       ValidationErrorCode,
		Message:    DefaultErrorMessage,
		Detail:     detail,
	}
}

func UnauthorizedError(detail string) HandledError {
	return HandledError{
		HttpStatus: http.StatusUnauthorized,
		Code:       UnauthorizedErrorCode,
		Message:    DefaultErrorMessage,
		Detail:     detail,
	}
}

func NotFoundError(detail string) HandledError {
	return HandledError{
		HttpStatus: http.StatusNotFound,
		Code:       NotFoundErrorCode,
		Message:    DefaultErrorMessage,
		Detail:     detail,
	}
}

func BadRequest(detail string) HandledError {
	return HandledError{
		HttpStatus: http.StatusBadRequest,
		Code:       BadRequestCode,
		Message:    DefaultErrorMessage,
		Detail:     detail,
	}
}
