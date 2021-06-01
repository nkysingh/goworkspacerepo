package main

import "fmt"

func main() {
	m := map[string]int{
		"nky":   25,
		"hobby": 32,
	}
	//for map
	m["vkm"] = 31
	for k, v := range m {
		fmt.Println(k, v)
	}
	// for array
	x := []int{1, 21, 23, 3}
	for i, v := range x {
		fmt.Println(i, v)
	}
}
