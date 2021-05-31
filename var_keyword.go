package main
import (
	"fmt"
	) 
var y = 44
var z int
func main(){
	x:= 42
	fmt.Println(x)

	fmt.Println(y)

	foo()
	fmt.Println(z)
}

func foo(){
	fmt.Println("again im here",y)
}