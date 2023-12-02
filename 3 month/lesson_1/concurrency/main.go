package main

import (
	"fmt"
	"time"
)

func welcome() {
	fmt.Println("welcome")
}

func A() {
	fmt.Println("A")
}

func B() {
	fmt.Println("B")
}

func C() {
	fmt.Println("C")
}

func main() {

	// goroutine
	go welcome()
	go A()
	go B()
	go C()

	fmt.Println("Hello")

	time.Sleep(time.Nanosecond * 1)
}
