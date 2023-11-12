package main

import (
	"fmt"
)

func main() {
	fmt.Println(sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))
	fmt.Println(sortPeople([]string{"IEO", "Sgizfdfrims", "QTASHKQ", "Vk", "RPJOFYZUBFSIYp", "EPCFFt", "VOYGWWNCf", "WSpmqvb"}, []int{17233, 32521, 14087, 42738, 46669, 65662, 43204, 8224}))
}
func sortPeople(names []string, heights []int) []string {
	fmt.Println(heights)
	fmt.Println(names)
	for i := 0; i < len(heights)-1; i++ {
		for j := 0; j < len(heights)-i-1; j++ {
			if heights[j] < heights[j+1] {
				heights[j], heights[j+1] = heights[j+1], heights[j]
				names[j], names[j+1] = names[j+1], names[j]
			}
		}
	}
	fmt.Println(names)
	fmt.Println(heights)
	return names
}
