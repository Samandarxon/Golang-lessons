package calc

import (
	"errors"
	"fmt"
)

func Sub(a, b int) int {
	return a - b
}

func Mult(a, b int) int {
	return a * b
}

func Div(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New(fmt.Sprintf("b %d -ga teng bo'lishi mumkin emas", b))

	}
	return float64(a / b), nil
}
