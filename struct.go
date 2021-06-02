package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretagent struct {
	person
	ltk bool
}

func main() {

	sa := secretagent{
		person: person{
			first: "nicky",
			last:  "chandel",
			age:   27,
		},
		ltk: true,
	}
	p1 := person{
		first: "diksha",
		last:  "Bisen",
		age:   26,
	}

	fmt.Println(p1.first, p1.age)
	fmt.Println(sa.person.first, sa.last, sa.age, sa.ltk)

	//anonymous struct

	ana := struct {
		sports string
		player int
	}{
		sports: "cricket",
		player: 11,
	}

	fmt.Println(ana.sports, ana.player)
}
