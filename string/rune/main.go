package main

import "fmt"

func main() {
	s := "Hellõ World"
	r := []rune(s)

	fmt.Println("len(r)", len(r))
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c ", r[i])
	}
	fmt.Println("")
	for i := 0; i < len(r); i++ {
		fmt.Printf("%v ", r[i])
	}
	fmt.Println("")
	for i := 0; i < len(r); i++ {
		fmt.Printf("%x ", r[i])
	}
	fmt.Println("")
	for i := 0; i < len(r); i++ {
		fmt.Printf("%T ", r[i])
	}
	fmt.Println("")
}
