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

const BASE_URL string = "http://localhost:8080"

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello Get")
	})

	//buat request
	request, _ := http.NewRequest(http.MethodGet, BASE_URL+"/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	if response.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		assert.Equal(t, "Hello Get", string(body))
	}
}
