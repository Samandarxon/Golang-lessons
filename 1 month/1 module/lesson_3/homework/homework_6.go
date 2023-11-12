// 6. Piramida
package main

import "fmt"

func main(){
	var n=30
	var str = ""
	for i:=1; i<=n;i++{
		for j:=n; j>=1;j--{
			if i>=j{
				str+=" *  "
			}else{
				str+="  "
			}
		}
		str+="\n"
	}
	fmt.Println(str)
}
