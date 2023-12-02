package main

import "fmt"

func main() {
	// var n = -32121. Shu raqamni reverse qling.
	// var result = -12132

	var (
		n      = 32121
		number int
	)
	fmt.Println(n)

	for n != 0 {
		number = (number + n%10) * 10
		n = (n - n%10) / 10
	}
	number /= 10
	fmt.Println(number)

}
