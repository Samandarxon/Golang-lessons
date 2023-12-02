package main
import "fmt"

func main(){

	// var s = “Hello World” shu string ichidagi unli harflarni 
	// neshtaligini sanag va unli harflarni orniga pustoy string
	// blan almashtring “”.
	var (
		s = "Hello World"
		str string
	)

	for i:=0 ; i<len(s);i++{
		switch string(s[i]){
		case "a","A","e","E","I","i","O","o","U","u":
			str+=""
		default:
			str+=string(s[i])
		}
	}
	fmt.Println(str)





}