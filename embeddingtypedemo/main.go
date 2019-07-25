package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) PrintName() {
	fmt.Println("Name:", p.Name)
}
func (p *Person) SetAge(age int) {
	p.Age = age
}

func (s *Singer) SetWork(works string) {
	s.works = works
}

type Singer struct {
	Person // extends Person by embedding it
	works  string
}

func main() {
	var singer = Singer{Person: Person{"John", 30}}
	singer.PrintName()
	singer.Name = "John Smit"
	(&singer).SetAge(31)
	(&singer).PrintName()
	fmt.Println(singer.Age)
	(&singer).SetWork("Music Company")
	fmt.Println(singer.works)
}
