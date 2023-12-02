package main

import "fmt"

func main(){
	var x,a,y float64

	fmt.Print("konfet massasi x=:")
	fmt.Scanln(&x)
	fmt.Printf("%f kg konfet narxini kiriting a =:",x)
	fmt.Scanln(&a)
	fmt.Print("Necha klogram konfet olishingizni kiriting a =:")
	fmt.Scanln(&y)
	fmt.Printf("1kg konfet %f so'm turadi %fkg konft %f so'm bo'ladi\n",a/x,y,a/x*y)
	
}