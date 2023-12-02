package main

import (
	"fmt"
	"math"
	"time"
)

func Prime(Channel chan string, num int) {
	defer close(Channel)
	var (
		isPrime = true
	)
	if num < 2 {
		Channel <- fmt.Sprintf("%d is not prime number", num)
		return
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		Channel <- fmt.Sprintf("%d is prime number", num)
	} else {
		Channel <- fmt.Sprintf("%d is not prime number", num)
	}
}
func main() {

	var (
		number   int
		BChannel = make(chan string, 1)
	)
	fmt.Print("Enter a number:")
	fmt.Scan(&number)
	x_T := time.Now()
	go Prime(BChannel, number)

	fmt.Println(<-BChannel, "\nTime: ", time.Since(x_T))
}
