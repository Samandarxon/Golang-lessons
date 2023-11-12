package main

import "fmt"
func Set(x_slice ...int)[]int{
	var arr []int=[]int{}
	// fmt.Printf("%T %#v\n",x_slice,x_slice)
	for i,el:=range x_slice{
		// fmt.Printf("%d   =>>>><<<<=   %t\n",el,include(x_slice,el))
		if include(x_slice[i:],el) {
			arr=append(arr,el)
		}
	}
	return arr
}

func include(nums []int,n int)bool{
	count:=0
	for _,el:=range nums{
		if el == n{
			count++
		}
	}
	// fmt.Println(n)
	if count<2{
		return true
	}
	return false

}
func removeDuplicates(nums []int) int {
    // expectedNums:=nums
    nums = append(nums[:0],Set(nums...)...)
    k:= len(nums)
	fmt.Println("nums",nums)
	return k
}


func main(){
	fmt.Println(removeDuplicates([]int{1,1,2}))
}