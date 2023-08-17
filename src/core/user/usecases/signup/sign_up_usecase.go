package signup

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/ortegasixto7/go-bank-api/src/core/user"
	"github.com/ortegasixto7/go-bank-api/src/external/auth"
	"golang.org/x/crypto/bcrypt"
)

func Execute(request *SignUpRequest, userPersistence user.IUserPersistence, authService auth.AuthService) {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)

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

	// fmt.Printf("%v \n", user)
	// fmt.Printf("%v \n", userPersistence)

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go userPersistence.Create(&user, &waitGroup)
	go authService.Create(&auth, &waitGroup)
	waitGroup.Wait()

	defer fmt.Println("User Saved")
	// fmt.Printf("%v \n", request)

}
