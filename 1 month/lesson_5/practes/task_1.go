package main

import "fmt"

func pow(a,b float64) float64{
	var result float64=1
	if b==0 {
		return 1
		}else if b<0{
			// fmt.Println(a,b)
			for b<0 {
			result*=a
			b++
		}
	}
	for b>0{ 
		result*=a
		b--
	}
	return result
}
func main(){
	fmt.Println(pow(3,2))
}