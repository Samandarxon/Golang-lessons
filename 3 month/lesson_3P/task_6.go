package main

import (
	"fmt"
	"time"
)

func Resever(Channel chan string, str string) {
	defer close(Channel)
	result := ""
	for _, v := range str {
		result = string(v) + result
	}
	Channel <- result

}
func main() {
	var (
		bch   = make(chan string, 1)
		X_Str = "Create a program that uses goroutines and channels to reverse a string concurrently"
	)
	x_time := time.Now()
	go Resever(bch, X_Str)

	fmt.Println("Max: ", <-bch, "\nx_time: ", time.Since(x_time))

}
