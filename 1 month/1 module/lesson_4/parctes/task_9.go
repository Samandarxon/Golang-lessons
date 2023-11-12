package main
import "fmt"

func main(){
	// Sizga s string berilgan uni palindrome palindrome masligini toping.
	//  Example 
// s = “aba” // Palindrome
// s = “aac” // Palindrome emas
var str string = "wwwww2eaaewwwww"
var is_str bool = false

for i:=len(str)-1 ; i>0; i--{
	if str[i]==str[len(str)-i-1]{
		is_str=true
	}else{
		is_str=false
		break
	}
}
if is_str{
	fmt.Println("Palindrome")
}else {
	fmt.Println("It's not Palindrome")
}

}