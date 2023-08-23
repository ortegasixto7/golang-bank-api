package signup

import (
	"github.com/ortegasixto7/go-bank-api/src/exception"
)

type SignUpRequest struct {
	FirstName string
	LastName  string
	UserName  string
	Password  string
}

func (this *SignUpRequest) Validate() error {
	if this.FirstName == "" {
		return exception.BadRequestError(string(exception.FIRST_NAME_IS_REQUIRED))
	}
	if this.LastName == "" {
		return exception.BadRequestError(string(exception.LAST_NAME_IS_REQUIRED))
	}
	if this.UserName == "" {
		return exception.BadRequestError(string(exception.USER_NAME_IS_REQUIRED))
	}
	if this.Password == "" {
		return exception.BadRequestError(string(exception.PASSWORD_IS_REQUIRED))
	}
	return nil
}
