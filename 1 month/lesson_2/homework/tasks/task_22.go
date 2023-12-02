package main

import "fmt"

func main() {
	var a, b float64

	fmt.Print("a=")
	fmt.Scanln(&a)
	fmt.Print("b=")
	fmt.Scanln(&b)
	a, b = b, a
	fmt.Printf("a=%v\nb=%v\n", a, b)
}
