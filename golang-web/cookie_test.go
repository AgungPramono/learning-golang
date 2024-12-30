package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "AG-001"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprintf(writer, "Success create cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("AG-001")
	if err != nil {
		fmt.Fprintf(writer, "No Cookie")
	} else {
		fmt.Fprintf(writer, "Hello %s", cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/?name=agung", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)
	cookies := recorder.Result().Cookies()
	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s \n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "AG-001"
	cookie.Value = "agung"
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, err := io.ReadAll(recorder.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
