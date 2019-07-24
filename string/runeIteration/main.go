package main

import "fmt"

func main() {
	s := "Hellõ World"

	for index, char := range s {
		fmt.Printf("character at index %d is %c\n", index, char)
	}
}
