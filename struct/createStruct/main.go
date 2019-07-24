package main

import "fmt"

type Employee struct {
	firstName, lastName string
	salary              int
	fullTime            bool
}

func main() {
	var ross Employee
	fmt.Println(ross)
}
