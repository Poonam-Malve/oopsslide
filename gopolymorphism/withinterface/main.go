package main

import "fmt"

//Part 1 OMIT
type Animal interface {
	makeNoise() string
}

type Dog struct {
	name string
	legs int
}

func (d *Dog) makeNoise() string {
	return d.name + " says woof!"
}

type Duck struct {
	name string
	legs int
}

//Part 2 OMIT
func (d *Duck) makeNoise() string {
	return d.name + " says quack!"
}

func NewDog(name string) Animal {
	return &Dog{
		legs: 4,
		name: name,
	}
}

func NewDuck(name string) Animal {
	return &Duck{
		legs: 4,
		name: name,
	}
}

func main() {
	var dog, duck Animal

	dog = NewDog("Tommy")
	duck = NewDuck("donald")

	fmt.Println(dog.makeNoise())
	// Tommy says woof!
	fmt.Println(duck.makeNoise())
	// donald says quack!
}

//END
