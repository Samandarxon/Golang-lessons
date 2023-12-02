package main

import (
	"fmt"
)

func main(){
	var (
		num    = 12357
		number = 0
		sum    = 0
		count  = 0
	  )
	
	  for num > 0 {
		count = 0
		number = num % 10
		for i := 2; i <= number; i++ {
		  if number%int(i) == 0 {
			count++
		  }
		}
		if count == 1 {
		  sum += number
		}
	
		num /= 10
	  }
	  fmt.Println(sum)
}