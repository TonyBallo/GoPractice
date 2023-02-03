package main 

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]unit8 {
	slice := make([][]uint8, dy) //making an empty 2d slice of length dy
	for y := range slice { // filling the first dimension with empty 1d slices of length dx
		res[y] = make([]uint8, dx) //making the empty 1d slices of length dx
		for x := range slice[y] { //filling each index of the 2d slice with uint8s
			slice[y][x] = uint8((x + y) / 2) //making the ints functions of the index's
		}
	}
	return slice //return completed slice
}

func main() {
	pic.Show(Pic)
}