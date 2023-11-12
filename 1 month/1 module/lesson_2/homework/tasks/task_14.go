package main
import "fmt"

func main(){
var (
	L = 3
	S float64
	R float64
)
const PI = 3.14

	R = float64(L)/(2*PI)
	S = PI * float64(R*R)
	fmt.Printf("R = %.2f\nS = %.2f\n",R, S)
}