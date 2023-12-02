// Sizga butun sonli array berilgan va n butun son beriladi siz 
// agarda shu masiv ichida k uniqe bolsa -1, agarda uniqe bo'lmasa qancha
// duplicat bolgani chiqarishingiz kerak.
// Masalan : [1,1,12,23,3,10] n = 12
// Output:-1
// [1,2,4,4,5,6,7,12,20] n = 4

package main

import "fmt"

func main(){
	var arr = []int{1,2,4,4,5,6,7,12,20}
	var n = 4
	var is int
	var x_map=make(map[int]int)
	// fmt.Println(x_map[n])
	for _,el:=range arr{
		x_map[el]++
	}
	for key,val:=range x_map{
		if val==1 && n==key{
			is=-1
			// fmt.Println("Vaalaykum assalom")
		}else if val>1 && key==n{
			// fmt.Println("salom")
			is=x_map[n]
		}
	} 
	fmt.Println(is)
}