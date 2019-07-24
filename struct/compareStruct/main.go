package main

import "fmt"

type Employee struct {
	firstName, lastName string
	salary              int
}

func main() {
	ross := Employee{
		firstName: "Ross",
		lastName:  "Geller",
		salary:    1200,
	}

	rossCopy := Employee{
		firstName: "Ross",
		lastName:  "Geller",
		salary:    1200,
	}

	fmt.Println(ross == rossCopy)
} //End
