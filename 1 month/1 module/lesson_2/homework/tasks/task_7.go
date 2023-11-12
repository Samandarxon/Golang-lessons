package main

import "fmt"

func main(){
	var (
		R= 4
		L float32
		S float32
	)
	const PI = 3.14

	S = PI * float32(R*R)
	L = 2 * PI * float32(R)

	fmt.Println("S =:",S,"\nL =:",L)
}