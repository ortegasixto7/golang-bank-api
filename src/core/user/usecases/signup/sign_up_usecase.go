package signup

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ortegasixto7/go-bank-api/src/core/user"
)

func Execute(request SignUpRequest, userPersistence user.IUserPersistence) {

	var user user.User
	user.Id = strconv.FormatInt(time.Now().UnixMilli(), 10)
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.UserName = request.UserName
	fmt.Println(user.Id)

	// fmt.Printf("%v \n", user)
	// fmt.Printf("%v \n", userPersistence)

	userPersistence.Create(&user)

	defer fmt.Println("User Saved")
	// fmt.Printf("%v \n", request)

}
