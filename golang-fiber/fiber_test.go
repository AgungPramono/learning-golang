package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestFormRequest(t *testing.T) {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		return c.SendString("hello " + name)
	})

	body := strings.NewReader("name=agung")
	request := httptest.NewRequest(http.MethodGet, "/hello", body)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "hello agung", string(bytes))
}

//go:embed source/contoh.txt
var contohFile []byte

func TestMultipartFile(t *testing.T) {
	app := fiber.New()
	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		err = c.SaveFile(file, "./target/"+file.Filename)
		if err != nil {
			return err
		}
		return c.SendString("Uploaded file Success")
	})

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, err := writer.CreateFormFile("file", "contoh.txt")
	assert.Nil(t, err)
	file.Write(contohFile)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	//
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
	//
	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Uploaded file Success", string(bytes))
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestRequestBody(t *testing.T) {
	app := fiber.New()
	app.Post("/login", func(c *fiber.Ctx) error {
		body := c.Body()

		request := new(LoginRequest)
		err := json.Unmarshal(body, &request)
		if err != nil {
			return err
		}
		return c.SendString("Hello " + request.Username)
	})

	body := strings.NewReader(`{"username":"agung","password":"123456"}`)
	request := httptest.NewRequest(http.MethodPost, "/login", body)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello agung", string(bytes))
}

type RegisterRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
	Name     string `json:"name" xml:"name" form:"name"`
}

var app = fiber.New()

func TestBodyParser(t *testing.T) {
	app.Post("/register", func(c *fiber.Ctx) error {
		request := new(RegisterRequest)
		err := c.BodyParser(request)
		if err != nil {
			return err
		}
		return c.SendString("Register Success " + request.Username)
	})
}

func TestBodyParserJson(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`{"username":"agung","password":"123456","name":"agung"}`)
	request := httptest.NewRequest(http.MethodPost, "/register", body)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success agung", string(bytes))
}

func TestBodyParserForm(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`username=agung&password=123456&name=agung`)
	request := httptest.NewRequest(http.MethodPost, "/register", body)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success agung", string(bytes))
}

func TestBodyParserXml(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`
		<RegisterUser>
			<username>agung</username>
			<password>123456</password>
			<name>agung</name>
		</RegisterUser>
	`)
	request := httptest.NewRequest(http.MethodPost, "/register", body)
	request.Header.Add("Content-Type", "application/xml")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success agung", string(bytes))
}

func TestResponseJSON(t *testing.T) {
	app.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"username": "agung",
			"password": "rahasia",
			"name":     "agung",
		})
	})

	request := httptest.NewRequest(http.MethodGet, "/user", nil)
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"name":"agung","password":"rahasia","username":"agung"}`, string(bytes))
}
