package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string] int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range words {
		_, ok := m[v]
		if ok { 	// if word is in map, increment
			m[v]++
			
		} else {	// else, set to 1
			m[v] = 1
		}
	}

	return m 	//return the generated map
}