package main

import (
	"fmt"

	"app/calculate"
)

func main() {

	fmt.Println(calculate.Add(10, 20))
	fmt.Println(calculate.Sub(10, 20))
	fmt.Println(calculate.Mult(10, 20))
	div ,_ := calculate.Div(10, 20)
	
	fmt.Println(div)
}
