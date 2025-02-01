package handlederror

import "errors"

type HandledError struct {
	HttpStatus int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Detail     string `json:"-"`
}

func (e HandledError) Error() string {
	return e.Detail
}

func (e HandledError) WithMessage(msg string) HandledError {
	e.Message = msg
	return e
}

func Extract(err error) HandledError {
	handledErr := HandledError{}
	isHandled := errors.As(err, &handledErr)
	if isHandled {
		return handledErr
	}

	return InternalServerError(err.Error())
}
