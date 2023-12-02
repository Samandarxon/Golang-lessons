package main

import "fmt"

func main(){
// var = “Hello wordl” string berilgan siz shundagi oxirgi soz chiqarishingiz kerak.
	var (
		str="Hello wordl"
		j int
		index int
		str_2 string
	)

	for i:=0 ; i<len(str); i++{
		if string(str[i]) == " "{
			j++
			index=i
		}
	}
	for  i:=index+1 ; i<len(str); i++{
		str_2+=string(str[i])
	}
	fmt.Println(str_2)
}