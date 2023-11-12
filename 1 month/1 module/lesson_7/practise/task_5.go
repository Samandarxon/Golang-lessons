package main

import "fmt"


func pow2(n int){
	var x = 1
	for x <= n{
		x*=2
		if x<=n{
			fmt.Println(x)
		}
	}
}
func main(){
	pow2(10)
}