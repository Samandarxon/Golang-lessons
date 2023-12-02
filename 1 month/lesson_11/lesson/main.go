package main

import "fmt"

func Add(a, b int, c *int) {
	*c = a + b
}
func Sub(a, b int, c *int) {
	*c = a - b
}
func Mult(a, b int, c *int) {
	*c = a * b
}
func Div(a, b int, c *int) {
	if b == 0 {
		fmt.Printf("%d ni %d ga bo'lib bo'lmaydi\n", a, b)
	} else {
		*c = a / b
	}
}

// Add(a, b, c int)
// Sub(a, b, c int)
// Mult(a, b, c int)
// Div(a, b, c int)

func main() {

	var a, b, c = 25, 5, 0
	// swap(&a, &b)

	fmt.Println(a, b)
	Add(a, b, &c)
	fmt.Println("Add=:", c)
	Sub(a, b, &c)
	fmt.Println("Sub=:", c)
	Mult(a, b, &c)
	fmt.Println("Mult=:", c)
	Div(a, b, &c)
	fmt.Println("Div=:", c)
}
