package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
    var count=0

	for _,el_1:=range jewels{
		for _,el_2:=range stones{
			if el_1==el_2{
				count++
			}
		}
	}
	return count
}

func main(){
	fmt.Println(numJewelsInStones("aA","aAAbbbb"))
}