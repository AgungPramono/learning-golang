package main

import "fmt"

type BandUserFilter func(name string) bool

func registerUserBanned(nameUser string, bandUserFilter BandUserFilter) {
	if bandUserFilter(nameUser) {
		fmt.Println("You Banned user:", nameUser)
	} else {
		fmt.Println("Welcome : ", nameUser)
	}
}

func main() {
	bannedUser := func(name string) bool {
		return name == "goblok"
	}
	registerUserBanned("doni", bannedUser)

	registerUserBanned("adi", func(name string) bool {
		return name == "adi"
	})
}
