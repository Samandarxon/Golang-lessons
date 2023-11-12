package main
import "fmt"
// 2. Write a program in C to read n number of values in an array and display them in reverse order.
// Test Data :
// Input the number of elements to store in the array :3
// Input 3 number of elements in the array :
// element - 0 : 2
// element - 1 : 5
// element - 2 : 7
// Expected Output :
// The values store into the array are :
// 2 5 7
// The values store into the array in reverse are :
// 7 5 2 

func main(){
	var n int
	fmt.Print("Array uzunligini kiriting n=")
	fmt.Scanln(&n)
	var arr =make([]int,n)
	var arr_2 []int
	fmt.Printf("%T\n",arr)

	for i:=0; i<n ; i++{
		fmt.Printf("%d elementni kiriting arr[%d]=",i,i)
		fmt.Scanln(&arr[i])
	}
	fmt.Println("The values store into the array are: \n",arr)
	for i,_:=range arr{
		// arr_2=append(arr_2,arr[len(arr)-i-1])
		fmt.Println(arr[len(arr)-i-1])
	}
	fmt.Println("The values store into the array in reverse are : \n",arr_2)

	
}