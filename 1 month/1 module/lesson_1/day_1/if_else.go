package main

import "fmt"

func main() {

	// logical operator: &&, ||, !
	var age uint8 = 80

	if age >= 18 && age <= 50 {
		fmt.Println("Disco welcome")
	} else if age > 50 {
		fmt.Println("Goodbye")
	} else {
		fmt.Println("Not")
	}
}




