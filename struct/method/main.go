package main

import (
	"fmt"
)

type Employee struct {
	firstName, lastName string
}

func main() {
	e := &Employee{"Ross", "Geller"}
	fmt.Println("Without receiver : ", getFullName(e.firstName, e.lastName))
	fmt.Println("With receiver : ", e.getFullName())
	e.changeFirstName("Monica")
	fmt.Println("Without pointer : ", e.getFullName())
	e.changeFirstNamePointer("Monica")
	fmt.Println("With pointer : ", e.getFullName())
}
func getFullName(firstName, lastName string) string {
	return firstName + " " + lastName
}
func (e Employee) getFullName() string {
	return e.firstName + " " + e.lastName
}
func (e Employee) changeFirstName(firstName string) {
	e.firstName = firstName
}
func (e *Employee) changeFirstNamePointer(firstName string) {
	e.firstName = firstName
} //End
