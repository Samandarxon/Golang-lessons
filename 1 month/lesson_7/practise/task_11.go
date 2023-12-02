package main

import "fmt"


func reverseString(s []byte)  {
	// fmt.Println(s)
	st:=make([]byte,len(s))
    for i,_:=range s{
		st[i]=s[len(s)-i-1]
		// s=append(s[i:],s[len(s)-i-1])
	}
	s=append(s[:0],st...)
	fmt.Println(string(s),st)
}


func main(){
var s = []byte{'h','e','l','l','o'}
// fmt.Println(map[int]int{1:2,3:4,5:7})
reverseString(s)
}