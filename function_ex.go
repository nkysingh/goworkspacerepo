package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type sagent struct {
	person
	ltk bool
}

func (s sagent) speak() {
	fmt.Println("i am", s.first, s.last, " the secceret agent")
}

func (p person) speak() {
	fmt.Println("i am", p.first, p.last, " the person speak.")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into brrrrr", h.(person).first)
	case sagent:
		fmt.Println("I was passed into brrrrr", h.(sagent).first)

	}
	fmt.Println("I was passed in bar", h)
}
func main() {
	s := sagent{
		person: person{
			"james",
			"bond",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}
	fmt.Println(s)
	s.speak()

	fmt.Println(p1)
	bar(p1)
	bar(s)
}
