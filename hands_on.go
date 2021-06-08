package main

import (
	"fmt"
)

type person struct { //create a type person struct

	first string
}

type human interface { //create a type human interface

	speak() //to implicitly implement the interface, a human must have the speak method

}

func (p *person) speak() { //attach a method speak to type person using a pointer receiver *person
	fmt.Println("hello i can speak")
}

func saySomething(h human) { //create func “saySomething”  have it take in a human as a parameter

	h.speak() //have it call the speak method

}

func main() {
	p1 := person{
		first: "vijendra",
	}
	saySomething(&p1) //you CAN pass a value of type *person into saySomething
	//you CANNOT pass a value of type person into saySomething as a parameter is pointer

	p1.speak()

}
