package golang_validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

// https://pkg.go.dev/github.com/go-playground/validator/v10#readme-baked-in-validations
func TestVaildator(t *testing.T) {
	validate := validator.New()
	if validate == nil {
		t.Error("validate is nil")
	}
}

func TestValidationVariabel(t *testing.T) {
	validate := validator.New()

	var user string = "agung"
	err := validate.Var(user, "required")
	if err != nil {
		t.Error(err)
	}
}

func TestValidationTwoVariabel(t *testing.T) {
	validate := validator.New()

	password := "rahasia"
	confirmPassword := "rahasia"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestMultiTagValidation(t *testing.T) {
	validate := validator.New()
	username := "agung"

	err := validate.Var(username, "required,alpha")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestMultiTagParameterValidation(t *testing.T) {
	validate := validator.New()
	username := "agung"

	err := validate.Var(username, "required,alpha,min=3,max=100")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestStructValidation(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "agung@mail.com",
		Password: "agung123",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationError(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "agu",
		Password: "ag",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		validationError := err.(validator.ValidationErrors)
		for _, fieldError := range validationError {
			fmt.Println("Error:", fieldError.Field(), "on tag:", fieldError.Tag(), "with error:", fieldError.Error())
		}
	}
}

func TestCrossFieldValidation(t *testing.T) {
	type RegisterPassword struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	validate := validator.New()
	request := RegisterPassword{
		Username:        "agung@mail.com",
		Password:        "agung123",
		ConfirmPassword: "agung123",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestNestedStructValidation(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Name    string  `validate:"required"`
		Age     int     `validate:"required"`
		Address Address `validate:"required"`
	}

	validate := validator.New()
	user := User{
		Name: "agung",
		Age:  20,
		Address: Address{
			City:    "Kediri",
			Country: "Indonesia",
		},
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}
