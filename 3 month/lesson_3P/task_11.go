package main

import (
	"fmt"
	"strconv"
	"time"
)

func Fibonacci(channel chan int, n int) {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	channel <- f[n]
}

func main() {
	var num = 10
	X_chan := make(chan int, 1+num*(1+num))
	x_time := time.Now()
	for i := 0; i <= num; i++ {
		go Fibonacci(X_chan, i)
		fmt.Print(strconv.Itoa(<-X_chan) + " ")
	}
	fmt.Println("\nTime: ", time.Since(x_time))

}
