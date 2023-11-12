package main
import "fmt"

// 3. Write a program in C to find the sum of all elements of the array.
// Test Data :
// Input the number of elements to be stored in the array :3
// Input 3 elements in the array :
// element - 0 : 2
// element - 1 : 5
// element - 2 : 8
// Expected Output :
// Sum of all elements stored in the array is : 15 

func main(){
	var n,m int
	fmt.Print("Array uzunligini kiriting n=")
	fmt.Scanln(&n)
	var arr =make([]int,n)
	fmt.Printf("%T\n",arr)

	for i:=0; i<n ; i++{
		fmt.Printf("%d elementni kiriting arr[%d]=",i,i)
		fmt.Scanln(&arr[i])
		m+=arr[i]
	}
	fmt.Println("Sum of all elements stored in the array is : ",m)
}