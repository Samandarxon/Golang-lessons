package main

import "fmt"

func main(){
	var (
		a =3
		b = 4
	)

	if a!=0 && b!=0{
		fmt.Printf("Yig'indisi a+b=: %d\nKo'paytmasi a * b =: %d\nKvadratlari a =:%d b =: %d\n",a+b,a*b,a*a,b*b)
	}
}