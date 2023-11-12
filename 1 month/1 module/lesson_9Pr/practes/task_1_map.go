package main

import "fmt"

func isPalindrom(slice []int)bool{
	is:=false
	for i:=0; i<len(slice); i++{
	// 	fmt.Printf(" slice[%d]==slice[%d]\n slice[%d]= %d slice[%d]=%d\n",
	// 		i,len(slice)-i-1,i,slice[i],len(slice)-i-1,slice[len(slice)-i-1],
	// )
		if slice[i]==slice[len(slice)-i-1]{
			is = true
		}else{
			return false
		}
	}
	return is
}

func main(){
	is := isPalindrom([]int{1,2,3,4,2,1})
	if is{
		fmt.Println("Palidrom")
	}else{
		fmt.Println("Palidrom emas")
	}
}