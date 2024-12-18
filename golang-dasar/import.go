package main

import (
	"fmt"
	"golang-dasar/database"
	_ "golang-dasar/database"
	"golang-dasar/helper"
)

func main() {
	result := helper.ViewMessage("Golang is Easy")
	fmt.Println(result)

	fmt.Println(helper.Application)
	//fmt.Println(helper.version) //version tidak bisa di akses dari luar

	connection := database.GetConnection()
	fmt.Println("connecting to database : ", connection)
}
