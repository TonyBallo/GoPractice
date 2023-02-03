package main

import (
	"fmt"
)

//Slices can be created with the built-in <make> functiuon; this is how you create
//dynamically-sized arrays

//The make function allocates a zeroed array and returns a slice that refers to that array


func main() {

	a := make([]int, 5) //len(a) = 5
	printSlice("a", a)

	//you can also specify the capcity of the slice 
	b := make([]int, 0, 5) //len(b)=0, cap(b)=5
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)


}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}


