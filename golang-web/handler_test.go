package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, reader *http.Request) {
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func TestServerMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, reader *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})

	mux.HandleFunc("/hai", func(writer http.ResponseWriter, reader *http.Request) {
		fmt.Fprint(writer, "Hai")
	})

	mux.HandleFunc("/images/", func(writer http.ResponseWriter, reader *http.Request) {
		fmt.Fprint(writer, "images")
	})

	mux.HandleFunc("/images/thumbnail", func(writer http.ResponseWriter, reader *http.Request) {
		fmt.Fprint(writer, "thumbnail")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
