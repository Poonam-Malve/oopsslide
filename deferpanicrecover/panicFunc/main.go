package main

import "fmt"

// access an element of a slice by index
func accessElement(a []int, index int) int {

	if len(a) > index {
		return a[index]
	} else {
		panic("Out of bound condition")
	}

}

func main() {
	a := []int{1, 2, 3}

	// access 3rd element
	fmt.Println(accessElement(a, 2))

	// access 4th element (out of bound)
	fmt.Println(accessElement(a, 3))
}
