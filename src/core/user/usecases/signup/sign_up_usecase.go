package signup

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/ortegasixto7/go-bank-api/src/core/user"
	"github.com/ortegasixto7/go-bank-api/src/external/auth"
)

func Execute(request SignUpRequest, userPersistence user.IUserPersistence, authService auth.AuthService) {

	var auth auth.Auth
	auth.Id = strconv.FormatInt(time.Now().UnixMilli(), 10)
	auth.UserName = request.UserName
	auth.Password = request.Password
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
