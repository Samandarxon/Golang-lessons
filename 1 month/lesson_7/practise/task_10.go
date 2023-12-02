package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	// fmt.Printf("%q\n",s)
	var str=""
	s=strings.TrimSpace(s)
	// fmt.Printf("%q\n",s)
    for i:=len(s)-1; i>=0; i--{
		if string(s[i])!=" "{
			str+=string(s[i])
		}else{
			break
		}
	}
	return len(str)
}

func main(){
	var s = "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}