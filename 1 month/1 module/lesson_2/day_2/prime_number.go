package main

import "fmt"

func main() {

	var (
		prime_count = 100
		number = 2
	)

	for prime_count > 0 {
		var is_primce bool = true
		for  i := 2; i < number; i++ {

			if number % i == 0 {
				is_primce = false
				break
			}
		}

		if is_primce {
			fmt.Printf("%d ",number)
			prime_count--
		}
		number++
	}

	fmt.Println()
}