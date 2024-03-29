package main

import (
	"fmt"
)

//Part 1 OMIT
type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct { //composition
	title   string
	content string
	author
}

//Part 2 OMIT
func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

type website struct { //Embedding
	posts []post
}

func (w website) contents() {
	fmt.Println("Contents of Website")
	fmt.Println()
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

//Part 3 OMIT
func main() {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post2 := post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	w := website{
		posts: []post{post1, post2},
	}
	w.contents()
}

//END OMIT
