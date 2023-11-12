package main

import "fmt"

func main() {
	// var str = “123ABDSCSL234t8” shu string da neshta son borligini toping. unicode.IsLetter, yokida unicode is.Digit method larini ishlatish taqiqlanadi.
	// Digit numbers: 0,1,2,3,4,5,6,7,8,9
	var (
		str          = "123ABDSCSL234t8"
		count_number int
		count_letter int
	)
	// 1. yo'l
	// for i := 0; i < len(str); i++ {
	// 	// fmt.Printf("%T  %s", str[i], str[i])
	// 	switch string(str[i]) {
	// 	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
	// 		count_number++
	// 	default:
	// 		count_letter++

	// 	}
	// }

	// 2. yo'l
	for i := 0; i < len(str); i++ {
		// fmt.Println(str[i])
		if str[i] > 48 && str[i] < 58 {
			count_number++
		} else {
			count_letter++
		}

	}
	fmt.Printf("Raqamlar soni := %d\nHarflar soni =: %d\n", count_number, count_letter)

}
