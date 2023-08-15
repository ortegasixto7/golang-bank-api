package main

import (
	"fmt"

	"github.com/ortegasixto7/go-bank-api/src/core/user/usecases/signup"
	"github.com/ortegasixto7/go-bank-api/src/persistence/postgres"
)

func main() {
	fmt.Println("\n\n\nHello World")
	var signUpRequest signup.SignUpRequest
	signUpRequest.FirstName = "Sixto"
	signUpRequest.LastName = "Ortega"
	signUpRequest.UserName = "ortegasixto7"
	signUpRequest.Password = "123456"
	postgres.Init()

	signup.Execute(signUpRequest, new(postgres.PostgresUserPersistence))

}
