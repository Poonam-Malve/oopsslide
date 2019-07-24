package main

import (
	"errors"
	"fmt"
)

func add(x, y int) (z int, e error) {
	defer func() {
		if r := recover(); r != nil {
			e = errors.New("Failed to add")
		}
	}()
	panic("HELP")
	return x + y, nil
}
func main() {
	sum, err := add(1, 2)
	if err != nil {
		fmt.Println("Error when adding number:", err)
		return
	}
	fmt.Println("Sum:", sum)
} //End
