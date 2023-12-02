package main

import "fmt"

func main() {
	var s = "hello world"
	// 	var s = “hello world”. Shu string reverse qlish kerak.
	// var result = “dlrow olleh”.
	var str string 
	for i := len(s)-1; i >=0; i-- {
		str += string(s[i])
	}
	fmt.Println(str)

}
