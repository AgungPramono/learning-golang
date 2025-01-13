package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
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

func TestHttpRequest(t *testing.T) {
	app := fiber.New()
	app.Get("/request", func(c *fiber.Ctx) error {
		firstName := c.Get("firstname")
		lastName := c.Cookies("lastname")
		return c.SendString("Hello " + firstName + " " + lastName)
	})

	request := httptest.NewRequest("GET", "/request", nil)
	request.Header.Set("firstname", "edi")
	request.AddCookie(&http.Cookie{Name: "lastname", Value: "supono"})

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello edi supono", string(body))
}

func TestRouteParameter(t *testing.T) {
	app := fiber.New()
	app.Get("/users/:userId/orders/:orderId", func(c *fiber.Ctx) error {

		userId := c.Params("userId")
		orderId := c.Params("orderId")
		return c.SendString("Get order from " + userId + " with order id " + orderId)
	})

	request := httptest.NewRequest("GET", "/users/agung/orders/0010", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Get order from agung with order id 0010", string(body))
}
