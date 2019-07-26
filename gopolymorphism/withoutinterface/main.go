package main

import "fmt"

type Animal struct {
	makeNoiseFn func(*Animal) string //attribute as function
	name        string
	legs        int
}

func (a *Animal) makeNoise() string { //just wrapper
	return a.makeNoiseFn(a)
}

func NewDog(name string) *Animal {
	return &Animal{
		makeNoiseFn: func(a *Animal) string {
			return a.name + " says woof!"
		},
		legs: 4,
		name: name,
	}
}

func NewDuck(name string) *Animal {
	return &Animal{
		makeNoiseFn: func(a *Animal) string {
			return a.name + " says quack!"
		},
		legs: 4,
		name: name,
	}
}

func main() {

	var dog, duck *Animal

	dog = NewDog("Tommy")
	duck = NewDuck("donald")

	fmt.Println(dog.makeNoise())
	// Tommy says woof!
	fmt.Println(duck.makeNoise())
	// donald says quack!
}
