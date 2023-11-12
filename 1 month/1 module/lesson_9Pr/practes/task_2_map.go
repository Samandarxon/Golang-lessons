package main

import "fmt"


func main(){
	slice := []int{1,1,4,5,3,3,3,3,1,2} 
	x_map := make(map[int]int)
	for _,el:=range slice{
		x_map[el]++
	}
	fmt.Println(x_map)
	max_key:=0
	max:=0
	for key,val:=range x_map{
		if max<val{
			max=val
			max_key=key
		}
	}
	fmt.Printf("Eng ko'p takrorlangan son %d %d marta takrorlangan\n",max_key,max)
}