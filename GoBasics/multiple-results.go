package main

import "fmt"

func math(x, y int) (int, int, int, int) {
	add := x + y 
	sub := x - y 
	mult := x * y 
	div := x / y 


	return add, sub, mult, div
}

func main() {
	a, b, c, d := math(5, 2)
	fmt.Println(a, b, c , d)
}