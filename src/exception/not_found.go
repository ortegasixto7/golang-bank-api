package exception

type NotFoundException struct {
	ErrorMessage string
}

func (this *NotFoundException) Error() string {
	return this.ErrorMessage
}

func NotFoundError(errorMessage string) error {
	return &NotFoundException{ErrorMessage: errorMessage}
}

type NotFound string

const (
	USER_NOT_FOUND NotFound = "USER_NOT_FOUND"
)
