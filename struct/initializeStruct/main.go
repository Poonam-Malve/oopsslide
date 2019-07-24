package main

import "fmt"

type Employee struct {
	firstName, lastName string
	salary              int
	fullTime            bool
}

func main() {
	ross := Employee{
		firstName: "ross",
		lastName:  "Bing",
		fullTime:  true,
		salary:    1200,
	}

	fmt.Println(ross)
}
