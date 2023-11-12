package main
import (
	"fmt"
	"strings"
)

func restoreString(s string, indices []int)string{
	fmt.Println(s)
	str:=make([]string,len(indices))
	for _,el:=range indices{
		
		str[indices[el]]+=string(s[el])
		// fmt.Printf("%s\n",string(s[el]))
	}
	return strings.Join(str,"")
}

func main(){
fmt.Println(restoreString("codeleet", []int{4,5,6,7,0,2,1,3}))
}