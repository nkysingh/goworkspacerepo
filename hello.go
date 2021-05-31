package main
import "fmt"
func main() {
	fmt.Println("hello")

	foo()

	n,_ := fmt.Println("somehting,434, random words")
	fmt.Println(n)
	for i:= 0; i < 20; i++{
		if i%2 == 0{
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("im foo")
}

