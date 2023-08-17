package main

import (
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/ortegasixto7/go-bank-api/src/controllers"
	"github.com/ortegasixto7/go-bank-api/src/persistence/postgres"
)

func main() {
	postgres.Init()
	var userController controllers.UserController
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		var data = make(map[string]string)
		data["message"] = "Hello world!!"
		return c.JSON(data)
	})

	app.Post("/users/sign-up/v1", userController.SignUp)

	app.Listen(":3000")

}
