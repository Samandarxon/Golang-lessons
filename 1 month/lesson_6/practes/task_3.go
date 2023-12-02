package main

import "fmt"

func main(){
	var m_slice = []int{23,45,1,2,4,67,8,4,5,99,5,6}
	var summ = 0

	for _,val:=range m_slice {
		summ+=val
	}
	fmt.Println(summ)
}