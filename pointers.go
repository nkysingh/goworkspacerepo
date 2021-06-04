package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "james bond",
	}
	fmt.Println(p1)
	changeme(&p1)
	fmt.Println(p1)

}

func changeme(p *person) {
	p.name = "vijendra singh"
	(*p).name = "nky singh"

}
