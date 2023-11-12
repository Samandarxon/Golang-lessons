package main

import "fmt"

func main() {
	// n 12 gacha bolgan son berilad Siz switch case ishlatgan holda
	// yilning qaysi fasligini aniqlashtirishingiz kerak bolad.
	var n = 8
	switch n {
	case 12, 1, 2:
		fmt.Println("Qish fasli")
	case 3, 4, 5:
		fmt.Println("Baxor fasli")
	case 6, 7, 8:
		fmt.Println("Yoz fasli")
	case 9, 10, 11:
		fmt.Println("Kuz fasli")
	default:
		fmt.Printf("%d bunaqa songa to'g'ri keladigan fasil yo'q")
	}
}
