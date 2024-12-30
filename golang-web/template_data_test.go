package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Belajar Golang Web",
		"Name":  "Golang Data Template",
		"Detail": map[string]interface{}{
			"Note": "Template data struct adalah kerangka atau cetakan untuk mendefinisikan struktur data dalam pemrograman",
		},
	})
}

func TestDataMapTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Detail struct {
	Note string
}
type Page struct {
	Title, Name string
	Detail      Detail
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Belajar Golang Web",
		Name:  "Golang Template Data Struct",
		Detail: Detail{
			Note: "Template data struct adalah kerangka atau cetakan untuk mendefinisikan struktur data dalam pemrograman",
		},
	})
}

func TestTemplateStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
