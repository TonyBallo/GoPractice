package main

import (
	"fmt"
	"math"
)

//Variables declared inside an 'if' short statment are also
//available inisde any of the 'else' blocks

// Both calls to pow() return their results before the call
// to fmt.Println() in main begins


func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { //can use v in this block
		return v 
	} else { // and this block
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
}