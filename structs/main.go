package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

type thing string

func main() {
	jim := person {
		firstName: "Jim",
		lastName: "Cooper",
		contactInfo: contactInfo{
			email: "jim@cooper.com",
			zipCode: 90210,
		},
	}

	jim.updateName("Christa")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (s *thing) updateString(t thing) {
	(*s) = t
}