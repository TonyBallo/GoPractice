package main

import "fmt"


// multiple defer statements execute in last-in-first-out
// like a stack

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

