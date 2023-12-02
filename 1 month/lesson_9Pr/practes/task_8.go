package main

import "fmt"
// type User struct{
// 	Name string
// }

// type Users []User


func smallerNumbersThanCurrent(nums []int) []int {
	var x_slice =[]int{}
    for _,el_1 :=range nums{
		count:=0
		for _,el_2 :=range nums{
			if el_1>el_2{
				count++
			}
		}
		x_slice=append(x_slice,count)

	}

	return x_slice
}

func main(){
	fmt.Println(smallerNumbersThanCurrent([]int{8,1,2,2,3}))
}