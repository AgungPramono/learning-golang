package main

import (
	_ "github.com/go-sql-driver/mysql"
	"golang-restful-api/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {

	//server := InitializedServer()
	//err := server.ListenAndServe()
	//helper.PanicIfErr(err)
}
