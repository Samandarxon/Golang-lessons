package main
import "fmt"

func main(){
	var (
		x1 = 2.3
		x2 = 1.2345
	)

	fmt.Println("|X1-X2| =",math.Abs(x1-x2))
}