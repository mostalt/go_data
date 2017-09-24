package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	kolya := person{firstName: "Kolyan", lastName: "Shirshov"}
	fmt.Println(kolya)
}
