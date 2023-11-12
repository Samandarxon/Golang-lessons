package main

import "fmt"

func Set(x ...int)[]int{
	var arr []int=[]int{}
	fmt.Printf("%T %#v\n",x,arr)
	for i:=0 ; i<len(x);i++{
		for j:=i ; j<len(x);j++{
			if x[i]!=x[j]{
				arr=append(arr,x[i])
			
				}else{
							
			}
		}
	}
	return arr
}
func main(){
	fmt.Println(Set([]int{1,2,2,4,5,4,3,6,7,8,7,6,9,6,7,8,9}...))
	// fmt.Println(Set([]int{1,2,2}...))
}
