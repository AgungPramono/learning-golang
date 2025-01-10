package main

import (
	"github.com/gofiber/fiber/v2"
	"golang-fiber/helper"
)

func main() {
	app := fiber.New()

	err := app.Listen("localhost:3000")
	helper.PanicIfError(err)
}
