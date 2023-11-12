package main

import "fmt"

var (
	m int
)

func main() {

	// var age float64 = 10
	// var first_number float64 = 12.4
	// fmt.Println(first_number + age)

	// var a, b, c = 10, 15, 20

	var (
		name string = "Jeck"
		age int = 23
		isMarried bool = true
		weight, height float64 = 67.4, 171.2
	)
	// var Age = 10
	// fmt.Println("age:", Age)

	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("isMarried:", isMarried)
	fmt.Println("width:", weight)
	fmt.Println("heigth:", height)
	fmt.Printf(" ")

	// fmt.Println("Hello world")
	// fmt.Printf("%d\n", 10)

	num, e, r := 10, 40, 6
	fmt.Println(num, e, r)
}
