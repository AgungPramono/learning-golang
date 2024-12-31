package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Execute handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("Afer Execute handler")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Execute")
		fmt.Fprint(writer, "Hello Middleware")
	})

	mux.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Test Handler Execute")
		fmt.Fprint(writer, "Hello Middleware")
	})

	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Panic Handler Execute")
		panic("ups,panic")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	server := http.Server{
		Handler: logMiddleware,
		Addr:    "localhost:8080",
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
