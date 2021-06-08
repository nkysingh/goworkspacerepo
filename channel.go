package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cr := make(<-chan int) //recieve
	cs := make(chan<- int) //send

	// c <- 5
	// c <- 43

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	//genral to specific converaton work
	fmt.Println("=================")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

	//specific to genral doesnt work
	// fmt.Printf("c\t%T\n", (chan int)(cr))
	// fmt.Printf("c\t%T\n", (chan int)(cs))

}
