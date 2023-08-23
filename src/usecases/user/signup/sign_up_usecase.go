package signup

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/ortegasixto7/go-bank-api/src/core/user"
	"github.com/ortegasixto7/go-bank-api/src/exception"
	"github.com/ortegasixto7/go-bank-api/src/external/auth"
	"golang.org/x/crypto/bcrypt"
)

func Execute(request *SignUpRequest, userPersistence user.UserPersistence, authService *auth.AuthService) error {
	requestErrorCode := request.Validate()
	if requestErrorCode != nil {
		return requestErrorCode
	}

	userExists := authService.GetByUserNameOrNil(request.UserName)
	if userExists != nil {
		return exception.BadRequestError(string(exception.UNAVAILABLE_USER_NAME))
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 2)

	var auth auth.Auth
	auth.Id = uuid.NewString()
	auth.UserName = request.UserName
	auth.Password = string(hashedPassword)
	auth.CreatedAt = time.Now().UnixMilli()

	var user user.User
	user.Id = auth.Id
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.UserName = request.UserName
	user.CreatedAt = auth.CreatedAt

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go userPersistence.Create(&user, &waitGroup)
	go authService.Create(&auth, &waitGroup)
	waitGroup.Wait()
	return nil
}
