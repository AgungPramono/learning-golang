package golang_validation

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

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
