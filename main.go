package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	kolya := person{
		firstName: "Kolya",
		lastName:  "Shirshov",
		contact: contactInfo{
			email:   "ololo@gmail.com",
			zipCode: 405405,
		},
	}

	kolya.updateName("Kolusha")
	kolya.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
