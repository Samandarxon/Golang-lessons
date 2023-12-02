package main
import "fmt"

func main(){
	var str = "TheBest,.?country!! in the World,.! Tashkent"
	var str_2 string
	for i:=0 ; i<len(str); i++{
		fmt.Println(string(str[i]))
		if string(str[i])==","||string(str[i])=="."||string(str[i])=="?"||string(str[i])=="!"||string(str[i])==" "{
			continue
		}else{
			str_2+=string(str[i])
		}
	}
	fmt.Println(str_2)
}

// var s = “TheBest,.?country!! in the World,.! Tashkent”
// output = “TheBestcountryintheWorldTashkent” for loop ishlatkan holda “,” “.” “?” “ “ “!” shu larni olip tashlagan holdan outputdaga javobni chiqarin. strings.Trim methodini ishlatin taqiqlanadi.

