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

func (d *Duck) makeNoise() string {
	return d.name + " says quack!"
}

//Part 2 OMIT
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
	fmt.Println(duck.makeNoise())

}

//END OMIT
