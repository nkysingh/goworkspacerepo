package main

import (
	"golang.org/x/crypto/bcrypt"

	"fmt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.Mincost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginpsswd := `password123`
	err = bcrypt.CompareHashPassword(bs, []byte(loginpsswd))
	if err != nil {
		fmt.Println("you can not login")
		return
	}
	fmt.Println("you are login")

}
