package main

import (
	"fmt"
)

func main() {
	too()
	foo("nky")
	s1 := woo("pune ?")
	fmt.Println(s1)
}

// basic function
func too() {
	fmt.Println("here is too function")
}

// function with argument
func foo(s string) {
	fmt.Println("hello", s)
}

// function which return
func woo(st string) string {
	return fmt.Sprint("hello where you live in, ", st)
}
