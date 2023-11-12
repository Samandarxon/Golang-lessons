package calculate


import (
	"fmt"
	"errors"
)
func Div(x,y int) (float64 ,error){
	if y == 0 {
		return -1.0, errors.New(fmt.Sprintf("%d shu sonni %d - ga bo'lib bo'lmaydi " ))
	}
	return float64(x)/float64(y), nil
}