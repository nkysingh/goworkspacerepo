package main

import (
	"fmt"
)

func main() {
	foo(2, 3, 4, 5, 6, 7, 8, 9)
	
}

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for the index number", i, "the sum of", v, "is", sum)

	}
	fmt.Println(sum)

}
