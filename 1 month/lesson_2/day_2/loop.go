package main

import "fmt"

func main() {

	// 1. C-style
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}

		fmt.Println(i)
	}

	// 2. Only condition
	i := float64(0)
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 3. Infinity loop

	i = 0
	for {
		fmt.Println("Infinity loop....")
		if i > 10 {
			break
		}
		i++
	}
}