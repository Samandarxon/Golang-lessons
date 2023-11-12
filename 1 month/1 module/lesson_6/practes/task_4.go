package main

import "fmt"

func main(){

var X_slice = []int{34,25,67,22,1,22,78,91,65,54,32,44}
var min = X_slice[0]
for _,val:=range X_slice{
	if min>val{
		min=val
	}

}
fmt.Println(min)
}