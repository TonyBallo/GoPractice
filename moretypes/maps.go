package main

import "fmt"

type Vertex struct {
	Lat, long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["place 2"] = Vertex{
		23.32143, -21.38549,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["place 2"])

}
