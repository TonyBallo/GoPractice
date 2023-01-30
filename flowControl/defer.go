package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world") //arguments are imediately evaluated, but the call is not executed until the surrounding function returns.

	fmt.Println("hello")
}