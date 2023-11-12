// Sizga butun sonli array berilgan sz uni ichigan duplicate bo'ldgan
// eng kotta butun sonni chop etishingiz kerak. Masalan
// [1,1,1,2,2,3,4,4,5,10]
// Output: 4
package main

import "fmt"

func main(){
	var arr = []int{1,1,1,2,2,3,4,4,5,10}
	var x_map = make(map[int]int)
	var max=0
	for _,el:=range arr{
		x_map[el]++
	} 
	fmt.Println(x_map)

	for key,val:=range x_map{
	// fmt.Println(key,val)

		if val==1{
			delete(x_map,key)
		}
		if _,ok:=x_map[key]; max<key && ok{
			max=key
		}
	}
	fmt.Println(max,x_map)
}