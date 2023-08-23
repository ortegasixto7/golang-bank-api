package exception

type BadRequestException struct {
	ErrorMessage string
}

func (this *BadRequestException) Error() string {
	return this.ErrorMessage
}

func BadRequestError(errorMessage string) error {
	return &BadRequestException{ErrorMessage: errorMessage}
}

type BadRequest string

const (
	FIRST_NAME_IS_REQUIRED BadRequest = "FIRST_NAME_IS_REQUIRED"
	LAST_NAME_IS_REQUIRED  BadRequest = "LAST_NAME_IS_REQUIRED"
	USER_NAME_IS_REQUIRED  BadRequest = "USER_NAME_IS_REQUIRED"
	PASSWORD_IS_REQUIRED   BadRequest = "PASSWORD_IS_REQUIRED"

	UNAVAILABLE_USER_NAME BadRequest = "UNAVAILABLE_USER_NAME"
)
