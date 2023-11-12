package main
import (
	"fmt"
	"math"
)

func main(){
	var x,y float64

	fmt.Print("x=")
	fmt.Scanln(&x)
	y = 4*math.Pow(x-3,2)- 7*math.Pow(x-3,3)+2
	fmt.Println("y =",y)
}