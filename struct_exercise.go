package main

import "fmt"

type person struct {
	first   string
	last    string
	fav_ice []string
}

func main() {
	p1 := person{
		first:   "diksha",
		last:    "bisen",
		fav_ice: []string{"vanila", "cold", "ice"},
	}
	fmt.Println(p1)
	for i, v := range p1.fav_ice {
		fmt.Println(i, v)
	}

	m := map[string]person{
		p1.last: p1,
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for l, b := range v.fav_ice {
			fmt.Println(l, b)

		}

	}
}
