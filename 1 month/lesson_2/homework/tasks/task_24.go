package main

import "fmt"

func main() {
	var a, b, c float64

	fmt.Print("a=")
	fmt.Scanln(&a)
	fmt.Print("b=")
	fmt.Scanln(&b)
	fmt.Print("c=")
	fmt.Scanln(&c)
	a, b, c = c, a, b
	fmt.Printf("                    \na=%v b=%v c=%v\n", a, b,c)
}
