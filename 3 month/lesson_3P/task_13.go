package main

import (
	"fmt"
	"strings"
	"time"
)

func StringMax(Channel chan string, str string) {
	defer close(Channel)
	max := 0
	maxStr := ""
	for _, v := range strings.Split(str, " ") {

		if max < len(v) {
			max = len(v)
			maxStr = v
		}
	}
	Channel <- maxStr

}
func main() {
	var (
		bch   = make(chan string, 1)
		X_Str = "Create a program that uses goroutines and channels to reverse a string concurrently"
	)
	x_time := time.Now()
	go StringMax(bch, X_Str)
	str := <-bch
	fmt.Println("Max Word: ", str, "Length: ", len(str), "\nx_time: ", time.Since(x_time))

}
