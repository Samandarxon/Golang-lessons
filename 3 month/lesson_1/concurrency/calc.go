package main

import (
	"fmt"
	"time"
)

func square(x chan int, input int) {
	time.Sleep(time.Second * 1)

	x <- input * input
}

func cube(y chan int, input int) {
	time.Sleep(time.Second * 2)

	y <- input * input * input
}

func main() {

	startTime := time.Now()

	var (
		squareInt = 3
		cubeInt   = 5

		squareChan = make(chan int)
		cubeChan   = make(chan int)
	)

	go square(squareChan, squareInt)
	go cube(cubeChan, cubeInt)

	fmt.Println("squareInt:", <-squareChan)
	fmt.Println("cubeInt:", <-cubeChan)

	fmt.Println("End time:", time.Since(startTime))
}
