package main

import "fmt"

func main() {
	monica := struct {
		firstName, lastName string
		salary              int
		fullTime            bool
	}{
		firstName: "Monica",
		lastName:  "Geller",
		salary:    1200,
	}

	fmt.Println(monica)
	fmt.Printf("%T", monica)
}
