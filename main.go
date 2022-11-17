package main

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id        string
	Firstname string
	Lastname  string
}

func handleUser(c *fiber.Ctx) error {
	user := User{
		Firstname: "Joe",
		Lastname:  "Doe",
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func handleCreateUser(c *fiber.Ctx) error {
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	userGroup := app.Group("/users")
	userGroup.Get("", handleUser)
	userGroup.Post("", handleUser)

	app.Listen(":3000")
}
