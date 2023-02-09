package main

import (
	"fmt"
	"math/rand"
)

func getSlot() int {
	return rand.Intn(3)
}

func main() {
	slots := []string {
		"|cherry|",
		"|bannana|",
		"|apple|",
	}

	fmt.Printf("\n\n\tWe Playin Slots Today!\n")
	fmt.Printf("\nPress <Enter> to continue play or <q> to quit\n")
	fmt.Println("###############################################")

	var keypress string

	fmt.Scanf("%s", &keypress)

	for ; keypress != "q"; {
		slot1, slot2, slot3 := getSlot(), getSlot(), getSlot()

		fmt.Println(slots[slot1], slots[slot2], slots[slot3])
		if slot1 == slot2 && slot2 == slot3 {
			
			fmt.Println("\n#######")
			fmt.Println("Winner!")
			fmt.Printf("#######\n")
		}
		fmt.Scanf("%s", &keypress)
	}
}