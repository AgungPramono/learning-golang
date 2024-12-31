package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotFoundHandler(t *testing.T) {

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Not Found")
	})

	//buat request
	request, _ := http.NewRequest(http.MethodGet, BaseUrl+"/404", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	if response.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		assert.Equal(t, "Not Found", string(body))
	}
}
