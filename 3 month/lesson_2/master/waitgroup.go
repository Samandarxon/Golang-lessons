package main

import (
	"fmt"
	"sync"
)

func A(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("A")
}

func B(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("B")
}

func C(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("C")
}

func main() {

	var wg = sync.WaitGroup{}

	go A(&wg)
	wg.Add(1)

	go B(&wg)
	wg.Add(1)

	go C(&wg)
	wg.Add(1)

	// wg.Add(3)

	wg.Wait()
}
