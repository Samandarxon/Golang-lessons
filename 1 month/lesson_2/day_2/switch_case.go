package main

import "fmt"

func main() {


	var (
		student_ball = 80
	)

	if 100 > student_ball && student_ball > 90 {
		fmt.Println("A")
	} else if 90 > student_ball && student_ball > 80 {
		fmt.Println("B")
	} else if 80 > student_ball && student_ball > 70 {
		fmt.Println("C")
	} else if 70 > student_ball && student_ball > 60 {
		fmt.Println("D")
	} else {
		fmt.Println("?")
	}


	var (
		traffic_light = "RED"
	)

	switch traffic_light {
	case "RED": 
		fmt.Println("STOP")
		fallthrough
	case "YELLOW": 
		fmt.Println("WARNING")
	case "GREEN": 
		fmt.Println("GO")
	default:
		fmt.Println("?")
	}

	// if condition {
	// } else if condition {
	// } else {
	// }

	// switch condition {
	// case element:
	// case element:
	// case element:
	// case element:
	// case element:
	// default: 
	// }
}
