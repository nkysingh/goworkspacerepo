package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "james",
		Last:  "bond",
		Age:   44,
	}

	p2 := person{
		First: "vijendra",
		Last:  "chandel",
		Age:   24,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
