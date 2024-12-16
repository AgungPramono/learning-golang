package main


import "fmt"


//named return values
//adalah sebuah fitur di Go yang memungkinkan kita untuk mengembalikan nilai lebih dari
//satu dari sebuah fungsi.
//tipe data parameter harus sama semua
func getCompleteName()(firstName, middleName,lastName string){
	firstName = "Agung"
	middleName = "Satria"
	lastName = "Budiono"
	return firstName,middleName,lastName
}

func main(){
	a,b,c := getCompleteName();
	fmt.Println(a,b,c)
}