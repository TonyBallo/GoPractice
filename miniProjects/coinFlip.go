package main

import (
	"fmt"
	"math/rand"
)

func main() {
	flips := []string {
		"Heads",
		"Tails",
	}

	fmt.Println("Lets flip a coin!")
	fmt.Println("Press <Enter> to flip or <q> to quit")

	var a string

	fmt.Scanf("%s", &a)

	for ; a != "q"; {
		fmt.Println(flips[rand.Intn(2)])
		fmt.Scanf("%s", &a)
	}
	



}