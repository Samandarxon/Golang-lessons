package main


import "fmt"

func div(x,y int) int{
	return x/y
}

func mult(x,y int)int{
	return x*y
}


func mod(x,y int)int{
	return x%y
}

func build(function func(int,int)int,x,y int) int{
	fmt.Println()
	return function(x,y)

}
// func build_2(function func(int,int)int,x,y int){
// 	fmt.Println(function(x,y))
// }
func main(){
	fmt.Println(build(div,40,2))
	fmt.Println(build(mult,40,2))
	fmt.Println(build(mod,40,9))

}