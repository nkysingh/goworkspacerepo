package main

import (
	"fmt"
)

func main() {
	a := struct {
		name string
		club map[string]int
		time []string
	}{
		name: "vijendra",
		club: map[string]int{
			"bar": 12,
			"MC":  21,
		},
		time: []string{
			"111",
			"333",
		},
	}
	fmt.Println(a.name)
	fmt.Println(a.club)
	fmt.Println(a.time)
	for k, v := range a.club {
		fmt.Println(k, v)
	}
	for k, v := range a.time {
		fmt.Println(k, v)
	}

}
