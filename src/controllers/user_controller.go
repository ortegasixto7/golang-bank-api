package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortegasixto7/go-bank-api/src/core/user"
	"github.com/ortegasixto7/go-bank-api/src/core/user/usecases/signup"
	"github.com/ortegasixto7/go-bank-api/src/external/auth"
	"github.com/ortegasixto7/go-bank-api/src/persistence/postgres"
)

type UserController struct {
	UserPersistence user.IUserPersistence
	AuthService     auth.AuthService
}

func (this *UserController) SignUp(c *fiber.Ctx) error {
	request := new(signup.SignUpRequest)

	if err := c.BodyParser(request); err != nil {
		return err
	}

	signup.Execute(request, new(postgres.PostgresUserPersistence),
		auth.AuthService{AuthPersistence: new(postgres.PostgresAuthPersistence)})

	c.Response().BodyWriter().Write([]byte(""))
	return nil
}
