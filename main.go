package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	var kolya person

	kolya.firstName = "Kolyan"
	kolya.lastName = "Shirshov"
	fmt.Println(kolya)
	fmt.Printf("%+v", kolya)
}
