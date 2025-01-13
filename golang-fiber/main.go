package main

import (
	"github.com/gofiber/fiber/v2"
	"golang-fiber/helper"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen("localhost:3000")
	helper.PanicIfError(err)
}
