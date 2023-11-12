package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
}

func checkIfPangram(sentence string) bool {
	map_x := make(map[string]int)
	count := 0
	for _, v := range sentence {
		map_x[string(v)]++
		if map_x[string(v)] == 1 {
			count++
		}
	}
	// fmt.Println(map_x)
	if count == 26 {
		return true
	}
	return false
}
