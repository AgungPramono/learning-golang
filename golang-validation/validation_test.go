package golang_validation

import (
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
