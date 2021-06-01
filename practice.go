package main

import "fmt"

func main() {
	s := []string{"vijendra", "chandel", "cricket", "bat", "reading", "books"}
	fmt.Println(s)
	v := []string{"tom", "moody", "loves", "swimming"}
	fmt.Println(v)
	new := [][]string{s, v}
	fmt.Println(new[1][3])

}
