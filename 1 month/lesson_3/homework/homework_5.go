package main

import "fmt"


func main(){
// 	5. Mukammal Sonni toping
// 6 = 1 + 2 + 3
// M: limit = 100
	var limit=100
	for i:=1;i<=limit;i++{
		sum:=0
		for j:=1;j<i;j++{
			if i%j==0{
				sum+=j
			}
		}
		if i==sum{
			fmt.Println(i)
		}
	}
	
}