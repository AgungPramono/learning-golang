package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
	"golang-fiber/helper"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Prefork:      true,
	})

	//middleware sederhana
	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("middleware before processing request")
		err := c.Next()
		fmt.Println("middleware after processing request")
		return err
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if fiber.IsChild() {
		fmt.Println("fiber is child process")
	} else {
		fmt.Println("fiber is parent process")
	}

	err := app.Listen("localhost:3000")
	helper.PanicIfError(err)
}
