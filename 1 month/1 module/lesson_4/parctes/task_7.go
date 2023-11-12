package main

import "fmt"

func main() {
	// var n = 5. for loop  ichida aylanip chiqip 2 dan tashqari
	//  hamma 5 gacha bolgan soni chiqarish kerak
	var n = 5
	for i := 0; i <= n; i++ {
		if i != 2 {
			fmt.Println(i)
		}
	}

}
