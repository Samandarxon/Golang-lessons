package main
import "fmt"

// 1. Write a program in C to store elements in an array and print them.
// Test Data :
// Input 10 elements in the array : 

func main(){
	var n int
	fmt.Print("Array uzunligini kiriting n=")
	fmt.Scanln(&n)
	var arr =make([]int,n)
	fmt.Printf("%T\n",arr)

	for i:=0; i<n ; i++{
		fmt.Printf("%d elementni kiriting arr[%d]=",i,i)
		fmt.Scanln(&arr[i])
	}
	fmt.Println(arr)
}