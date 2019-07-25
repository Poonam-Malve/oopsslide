package main

import "fmt"

//Part 1 OMIT

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

//Part 2 OMIT
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

//END OMIT
