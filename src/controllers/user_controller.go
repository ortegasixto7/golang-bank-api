package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortegasixto7/go-bank-api/src/core/user"
	"github.com/ortegasixto7/go-bank-api/src/external/auth"
	"github.com/ortegasixto7/go-bank-api/src/persistence/postgres"
	"github.com/ortegasixto7/go-bank-api/src/usecases/user/signup"
)

type UserController struct {
	UserPersistence *user.UserPersistence
	AuthService     *auth.AuthService
}

func (this *UserController) SignUp(c *fiber.Ctx) error {
	request := new(signup.SignUpRequest)

	if err := c.BodyParser(request); err != nil {
		return err
	}

	err := signup.Execute(request, new(postgres.PostgresUserPersistence), &auth.AuthService{AuthPersistence: new(postgres.PostgresAuthPersistence)})
	if err != nil {
		return err
	}

	c.Response().BodyWriter().Write([]byte(""))
	return nil
}
