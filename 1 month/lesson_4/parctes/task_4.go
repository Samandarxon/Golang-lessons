package main

import "fmt"

func main(){
	// n 7 gacha bolgan son berilad. Siz switch case 
	// ishlatgan holda haftaning qaysi kuniligini aniqlashtirishingiz 
	// kerak bolad.
	var n int =4

	switch n {
	case 1:
		fmt.Println("Dushanba")
	case 2:
		fmt.Println("Sesanba")
	case 3:
		fmt.Println("Chorshanba")
	case 4:
		fmt.Println("Payshanba")
	case 5:
		fmt.Println("Juma")
	case 6:
		fmt.Println("Shanba")
	case 7:
		fmt.Println("Yakshanba")
	default:
		fmt.Println("Bunaqa kun yo'q")
	} 

}