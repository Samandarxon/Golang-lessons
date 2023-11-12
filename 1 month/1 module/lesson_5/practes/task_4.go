package main

import "fmt"

func Floor(a float64) int {
		if a>0 {
			return int(a)
		}
	return int(a)-1		
}

func main() {
	fmt.Println(Floor(-18.8))

}
