package main

import (
	"fmt"
)

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func FullName(firstName *string, lastName *string) {
	defer recoverName()
	if firstName == nil {
		panic("firstname can not be nill")
	}
	if lastName == nil {
		panic("last name can not be nill")
	}
	fmt.Println(*firstName, *lastName)
	fmt.Println("returned noramlly from function Fullname")
}

func main() {
	defer fmt.Println("defered call from main ")
	firstName := "james"
	FullName(&firstName, nil)
	fmt.Println("Return norrmally from main")
}

// Recover works only when it is called from the same goroutine which is panicking.
// It's not possible to recover from a panic that has happened in a different goroutine.
