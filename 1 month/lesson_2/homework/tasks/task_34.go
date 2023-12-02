package main

import "fmt"

func main(){
	var x,a,y,b float64

	fmt.Print("Shokolad massasini kiriting x=:")
	fmt.Scanln(&x)
	fmt.Printf("%f kg shokolad narxini kiriting a =:",x)
	fmt.Scanln(&a)
	fmt.Print("Konfet massasini kiriting y=:")
	fmt.Scanln(&y)
	fmt.Printf("%f kg Konfet narxini kiriting a =:",y)
	fmt.Scanln(&b)
	fmt.Printf("1kg shokolad 1kg konfetdan %f so'm qimmat turadi\n",a/x-b/y)
	
}