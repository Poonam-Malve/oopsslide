package main

import "fmt"

func defBar() {
	fmt.Println("defBar() started")
	if r := recover(); r != nil {
		fmt.Println("WOHA! Program is panicking with value", r)
	}
	fmt.Println("defBar() done")
}
func defFoo() {
	fmt.Println("defFoo() Started")
	defer defBar() // defer defBar call
	fmt.Println("defFoo() done")
}
func normMain() {
	fmt.Println("normMain() started")
	defer defFoo() // defer defFoo call
	panic("HELP")  // panic here
	fmt.Println("normMain() done")
}
func main() {
	fmt.Println("main() started")
	normMain() // normal call
	fmt.Println("main() done")
} //End
