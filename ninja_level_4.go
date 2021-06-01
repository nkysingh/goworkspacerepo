package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken", "not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	xx := [][]string{x, y}

	for i, v := range xx {
		fmt.Println("record: ", i)
		for j, val := range v {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}

	}

}
