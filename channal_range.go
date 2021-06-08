package main

import "fmt"

// Method 1 start
func main() {
	c := make(chan int)

	go foo(c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("existsssssssssssssssssss")
}

func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

// Method 1 end

// Method 2 start

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("exsitssssssssssss")
}

// Method 2 endddddddddddddddddd
