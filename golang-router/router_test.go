package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

const BaseUrl string = "http://localhost:8080"

var router = httprouter.New()

func TestRouter(t *testing.T) {
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello Get")
	})

	//buat request
	request, _ := http.NewRequest(http.MethodGet, BaseUrl+"/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	if response.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		assert.Equal(t, "Hello Get", string(body))
	}
}
