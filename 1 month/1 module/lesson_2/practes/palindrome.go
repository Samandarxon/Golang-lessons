package main

import "fmt"

func main() {

	var (
		number = 123454321
		summer = 0
		num_clone = number
	)

	for number > 0 {
		summer = summer * 10 + (number % 10)
		number /= 10
	}

	if summer == num_clone {
		fmt.Println("Palindrome")
		return
	} 

	fmt.Println("Not Palindrome")
}