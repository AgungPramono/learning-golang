package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestRoutingHelloWorld(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	request := httptest.NewRequest("GET", "/", nil)
	reponse, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, reponse.StatusCode)

	body, err := io.ReadAll(reponse.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello, World!", string(body))
}
func TestContextQueryParam(t *testing.T) {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Query("name", "guest")
		return c.SendString("Hello " + name)
	})

	request := httptest.NewRequest("GET", "/hello?name=Agung", nil)
	reponse, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, reponse.StatusCode)

	body, err := io.ReadAll(reponse.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Agung", string(body))

	request = httptest.NewRequest("GET", "/hello?name=", nil)
	reponse, err = app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, reponse.StatusCode)

	body, err = io.ReadAll(reponse.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello guest", string(body))
}
