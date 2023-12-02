package main

import "fmt"

func main() {

	// Total
	var num_slice = []int{1, 2, 3, 4}
	fmt.Println("Length:", len(num_slice), "Capacity:", cap(num_slice))

	num_slice = append(num_slice, 10)
	fmt.Println("Length:", len(num_slice), "Capacity:", cap(num_slice))

	num_slice = append(num_slice, 20)
	fmt.Println("Length:", len(num_slice), "Capacity:", cap(num_slice))

	num_slice = append(num_slice, 30)
	fmt.Println("Length:", len(num_slice), "Capacity:", cap(num_slice))

	num_slice = append(num_slice, 40)
	fmt.Println("Length:", len(num_slice), "Capacity:", cap(num_slice))

	// for _, val := range num_slice {
	// 	fmt.Println(val)
	// }

	n := []int{11, 22, 33, 44, 55, 66}

	ind := 2

	if len(n) > ind {
		n = append(n[:ind], n[ind+1:]...)
		fmt.Println(n)
	} else {
		fmt.Println("no find element: ", ind)
	}

	// [:] // start, end

	// l := []int{77, 88, 99}

	// fmt.Println(n[3:4])
	// n = append(n, l...)
	// fmt.Println(n)
}
