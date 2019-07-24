package main

import "fmt"

type Salary struct {
	basic     int
	insurance int
	allowance int
}

type Employee struct {
	firstName, lastName string
	Salary
}

func main() {
	ross := Employee{
		firstName: "Ross",
		lastName:  "Geller",
		Salary:    Salary{1100, 50, 50},
	}

	ross.basic = 1200
	ross.insurance = 0
	ross.allowance = 0
	fmt.Println("Ross's basic salary", ross.basic)
	fmt.Println("Ross is", ross)
} //End
