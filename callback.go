package main

import "fmt"

func main() {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(i...)
	fmt.Println(s)
	ii := even(sum, i...)
	fmt.Println(ii)

}

func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)

}
