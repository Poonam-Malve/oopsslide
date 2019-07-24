package main

import (
	"fmt"
	"math"
)

//Part 1 OMIT
type geometry interface { //interface (Duck)
	area() float64
	perim() float64
}

type rectangle struct { // struct (Animal)
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 { //method
	return r.height * r.width
}
func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

//Part 2 OMIT
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func calculate(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main() {
	r := rectangle{width: 7, height: 9}
	c := circle{radius: 8}

	calculate(r)
	calculate(c)
}

//END OMIT
