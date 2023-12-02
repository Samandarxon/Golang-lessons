package main

import (
	"fmt"
	"strings"
)


func mostWordsFoud(sentences []string) int{
	max:=0
	for _,str_el:=range sentences{
		sentences_length := len(strings.Split(str_el," "))
		// fmt.Println("sentences_length = ",sentences_length)
		// fmt.Println(strings.Split(str_el," "))
		if max<sentences_length{
			max = sentences_length
		}
	}
	return max
}

func main(){

fmt.Println(mostWordsFoud([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}))


}