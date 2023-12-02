package main

import "fmt"

func main() {
	// Arifmetik Operator: +, -, /, *, %
	var (
		a int = 10
		b float32 = 22.23
		c float64 = 23.234
	)

	fmt.Println(a + int(b))
	fmt.Println(b + float32(c))



	var age uint8 = 255
	fmt.Println(age)

	// int8, int16, int32, int64, int
	// uint8, uint16, uint32, uint64, uint
}