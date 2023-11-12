package main

import "fmt"

func search(nums []int, target int) int {
    for i,el:=range nums{
        if el == target{
            return i
        }
    }
            return -1
}
func main(){
	fmt.Println(search([]int{-1,0,3,5,9,12},9))
}