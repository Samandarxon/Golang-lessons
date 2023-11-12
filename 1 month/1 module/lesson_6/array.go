package main

import "fmt"

func main() {

	// var numbers [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// var nums = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// var num_list = [...]int{1, 2, 3, 10}

	// fmt.Println("numbers length: ", len(numbers))
	// fmt.Println("nums length: ", len(nums))
	// fmt.Println("num_list length: ", len(num_list))

	var numbers [7]int
	numbers[0] = 10
	numbers[1] = 11
	numbers[2] = 12
	numbers[3] = 13
	numbers[4] = 14
	numbers[5] = 15
	numbers[6] = 16

	// numbers[11] = 23 ->> error
	// fmt.Println(numbers[0])

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Printf("%d ", numbers[i])
	// }
	// fmt.Println()

	for index, value := range numbers {
		fmt.Println(index, value)
	}
}
