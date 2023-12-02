package main

import "fmt"

const (
	ADD  = "+"
	SUB  = "-"
	MULT = "*"
	DIV  = "/"
)

func Calculator(a, b int) func(string) {
	c := 0
	return func(s string) {
		switch s {
		case "+":
			c = a + b
			fmt.Println(c)
		case "-":
			c = a - b
			fmt.Println(c)
		case "*":
			c = a * b
			fmt.Println(c)
		case "/":
			c = a / b
			fmt.Println(c)
		default:
			fmt.Printf("%s bunaqa matematik amal yo'q\n",s)
		}
	}
}

func main() {
	calculator:=Calculator(20,10)
	calculator(ADD)
	calculator(SUB)
	calculator(MULT)
	calculator(DIV)
	calculator("?")
}
