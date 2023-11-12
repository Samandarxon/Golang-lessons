package main

import "fmt"

func main(){
	var (
		d int = 5
		L float32
	)
	const PI = 3.14

	L= PI * float32(d)
	fmt.Printf("Aylana uzunligi L=:%f\n",L)

}