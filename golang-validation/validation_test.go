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
