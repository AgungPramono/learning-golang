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

func TestRouterPatternNamedParameter(t *testing.T) {
	router.GET("/product/:id/items/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		itemId := params.ByName("itemId")
		text := "Product " + id + " Item " + itemId
		fmt.Fprint(writer, text)
	})

	//buat request
	request, _ := http.NewRequest(http.MethodGet, BaseUrl+"/product/1/items/3", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	if response.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		assert.Equal(t, "Product 1 Item 3", string(body))
	}
}
