package main

import "fmt"

func main() {
	map_var := map[string]int{
		"age":   21,
		"class": 12,
	}
	fmt.Println(map_var)

	delete(map_var, "class")
	fmt.Println(map_var)

	delete(map_var, "age")
	fmt.Println(map_var)

}
