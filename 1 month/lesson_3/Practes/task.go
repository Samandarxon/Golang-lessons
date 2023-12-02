package main

import (
	"fmt"
)

func main(){
	// task_3(4)
	// task_4(1238)
	// task_5(12212)
	task_6(5)
}


func task_3(n int ){
	sum:=1
	for i:=2 ; i<=n ; i++{
		sum*=i
	}
	fmt.Println(sum)
}

func task_4(n int){
	sum:=0
	for n>0{
		sum += n%10
		n/=10
	}
	fmt.Println(sum)
}

func task_5(n int ){
	d := n
	number:=0
	for n>0{
			number= (number+n%10)*10
		n/=10
	}
	number = number/10
	
	if number == d{
		
		fmt.Println("Palidrom")
	}else {
		fmt.Println("Not Palindrome")
	}
}

func task_6(n int ){
	str:=""
	for i:=1; i<=n;i++{
		for j:=1; j<=i; j++{
			// if j==5{
				str+="*"
			// }else {
			// 	str+=" "
			// }
		}
		str+="\n"
	} 
	fmt.Println(str)
}