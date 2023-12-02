package main

import "fmt"

func main() {

	// var num int = 10
	// // var a int

	// // num = "ten"

	// fmt.Println(num)
	// fmt.Printf("%T\n", num)

	// int -> deafult: 0
	var a, b, c int = 10, 12, 14
	fmt.Println(a, b, c)

	// float32 va float64 -> deafult: 0
	var (
		d float32 = 12.24
		e float64 = 54.4
	)
	fmt.Println("Float->", d, e)

	// boolean -> deafult: false
	var isMarried bool = true

	fmt.Println("Boolean->", isMarried)
	// string -> deafult: ""
	var username string = "admin"
	fmt.Println("String->", username)
}