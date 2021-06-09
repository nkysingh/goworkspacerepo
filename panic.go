package main

import "fmt"

func slicePanic() {
	n := []int{23, 43, 45, 65}
	fmt.Println(n[4])
	fmt.Println("noramlly return from a")
}
func main() {
	slicePanic()
	fmt.Println("normally return the function")
}
