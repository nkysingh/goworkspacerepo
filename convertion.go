package main

import (
	"fmt"
)

var x int

type nky int

var y nky

func main() {
	x = 54
	y = 90
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	x = int(y)
	fmt.Println(x)
	fmt.Printf("%T\n", x)

}
