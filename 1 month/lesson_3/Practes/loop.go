package main

import(
	"fmt"
	"math"
)

func main(){
	// task_1(23)
	task_2(5)
}

func task_1(n int){
	i := 0
	for  n>0 {
		 n = (n-n%10)/10
		 i++
	}
	fmt.Println(i)
}

func task_2(n int){
	for d:=true, i:=2;n>0;i++ {
		
		for j:=2 ;j<=int(math.Sqrt(float64(i))) ; j++{
			if i%j == 0{
				d = false
				break
			}
		}
		if d{
			fmt.Println(i)
			n--
		}
	}
}