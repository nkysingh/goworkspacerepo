package main

import "fmt"

func main() {
	n, err := fmt.Print("hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
