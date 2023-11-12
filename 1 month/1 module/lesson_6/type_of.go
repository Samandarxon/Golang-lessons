package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num = 10
	var k = "10"

	fmt.Printf("%v: %T\n", num, num)

	fmt.Println(reflect.TypeOf(num))
	fmt.Println(reflect.TypeOf(num).Name() == "int")
	fmt.Println(reflect.TypeOf(num) == reflect.TypeOf(k))
}
