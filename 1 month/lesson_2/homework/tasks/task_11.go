package main
import (
	"fmt"
	"math"
)

func main(){
	var (
		a =3
		b = 4
	)

	if a!=0 && b!=0{
		fmt.Printf("Yig'indisi a+b=: %d\nKo'paytmasi a * b =: %d\nModullari a =:%v b =: %v\n",a+b,a*b,math.Abs(float64(a)),math.Abs(float64(b)))
	}	
}