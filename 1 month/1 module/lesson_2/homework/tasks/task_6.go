package main
import "fmt"

func main(){
	var (
		a =4
		b=2
		c=3
		V int
		S int
	)

	V=a*b*c
	S = 2*(a*b+b*c+c*a)
	fmt.Printf("Paralellopepid hajmi V=:%d\nParalellopepid to'la sirti yuzasi S=:%d\n",V,S)

}