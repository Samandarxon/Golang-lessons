package main

import (
	"errors"
	"fmt"
	"log"
)

// type Errlog struct {
// 	message string
// }

// func (e Errlog) Error() string {
// 	return e.message
// }

func divide(a, b int) (result int, err error) {

	if b == 0 {
		return 0, errors.New("non zero divide")
	}

	return a / b, nil
}

func main() {

	result, err := divide(12, 0)

	// err = nil -> err.Error()
	// err != nil -> err.Error()

	if err != nil {
		log.Println("Divide function error >>>>>> " + err.Error())
		return
	}

	fmt.Println("Result divide(3, 4) = ", result)
}
