package golang_validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"strings"
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

func TestCollectionStructValidation(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Name      string    `validate:"required"`
		Age       int       `validate:"required"`
		Addresses []Address `validate:"required,dive"`
	}

	validate := validator.New()
	user := User{
		Name: "",
		Age:  20,
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func TestBasicCollectionStructValidation(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Name      string    `validate:"required"`
		Age       int       `validate:"required"`
		Addresses []Address `validate:"required,dive"`
		Hobbies   []string  `validate:"required,dive,required,min=3"`
	}

	validate := validator.New()
	user := User{
		Name: "Agung",
		Age:  20,
		Addresses: []Address{
			{
				City:    "Surabaya",
				Country: "Indonesia",
			},
			{
				City:    "Kediri",
				Country: "Indonesia",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"X",
		},
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMapStructValidation(t *testing.T) {
	type School struct {
		Name string `validate:"required,min=2"`
	}

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Name      string            `validate:"required"`
		Age       int               `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=3"`
		Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys,required"`
	}

	validate := validator.New()
	user := User{
		Name: "Agung",
		Age:  20,
		Addresses: []Address{
			{
				City:    "Surabaya",
				Country: "Indonesia",
			},
			{
				City:    "Kediri",
				Country: "Indonesia",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"Fishing",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SD Negeri",
			},
			"SMP": {
				Name: "SMP Negeri",
			},
			"SMA": {
				Name: "i",
			},
		},
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestBasicMapValidation(t *testing.T) {
	type School struct {
		Name string `validate:"required,min=3"`
	}

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Name      string            `validate:"required"`
		Age       int               `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=3"`
		Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys,required"`
		Wallet    map[string]int    `validate:"dive,keys,required,endkeys,required,gt=1000"`
	}

	validate := validator.New()
	user := User{
		Name: "Agung",
		Age:  20,
		Addresses: []Address{
			{
				City:    "Surabaya",
				Country: "Indonesia",
			},
			{
				City:    "Kediri",
				Country: "Indonesia",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"Fishing",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SD Negeri",
			},
			"SMP": {
				Name: "SMP Negeri",
			},
			"SMA": {
				Name: "SM",
			},
		},
		Wallet: map[string]int{
			"Gopay":   2000000,
			"Mandiri": 3000,
			"Tunai":   10001,
		},
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestAliasTag(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	type Customer struct {
		Id    string `validate:"varchar,min=5"`
		Name  string `validate:"varchar"`
		Owner string `validate:"varchar"`
		Note  string `validate:"varchar"`
	}

	customer := Customer{
		Id:    "123",
		Name:  "agung",
		Owner: "agung",
		Note:  "",
	}

	err := validate.Struct(customer)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}
	return true
}

func TestMustValidUsername(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)

	type Customer struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	customer := Customer{
		Username: "AGUNG",
		Password: "1452556",
	}

	err := validate.Struct(customer)
	if err != nil {
		fmt.Println(err.Error())
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err.Error())
	}

	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}
	return len(value) == length
}

func TestValidPin(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidPin)

	type Login struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}

	login := Login{
		Phone: "123456",
		Pin:   "123456",
	}

	err := validate.Struct(login)
	if err != nil {
		fmt.Println(err.Error())
	}
}
