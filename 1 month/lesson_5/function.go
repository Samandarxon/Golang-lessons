package main


import "fmt"



// Argument function
func  add(a,b int){
	fmt.Printf("%d + %D = %d\n",a,b,a+b)
}

// Return function
func divide(a,b int)(int,bool){
	if b==0{
		return 0,true
	}
	return a/b ,false
}

func sqrt(x int )(square int){
	square = x * x
	return
}

func Color()(uint8,uint8,uint8){
	return 255, 255, 255
}

func main(){
	// welcom()
	// add(10,20)
	// result,mistake := divide(20,0)
	// if mistake{
	// 	fmt.Pirntln("Error")
	// }else{
	// 	fmt.Println(result)
	// }

	// fmt.Println(sqrt(10))

	// R,G,_:=Color()
	// fmt.Println(R,G)

	var R,G uint8 = 10, 10
	R,G,B:=Color()

	fmt.Println(R,G,B)
}