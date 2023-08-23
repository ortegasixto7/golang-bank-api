package main

import (
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-json"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/ortegasixto7/go-bank-api/src/controllers"
	"github.com/ortegasixto7/go-bank-api/src/persistence/postgres"
)

func main() {
	var PORT string = ":8005"
	var GO_ENV string = os.Getenv("GO_ENV")
	if GO_ENV == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		PORT = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}

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

	app.Listen(PORT)

}
