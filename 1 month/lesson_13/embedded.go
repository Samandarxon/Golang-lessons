package main

import "fmt"

// component

// getSugar()
// getName()

// cake
// candle

type Component struct {
	weight, sugar float64
}

type Cake struct {
	Component
	name string
}

type Candle struct {
	Component
	name string
}

func (c Component) getSugar() string {
	return fmt.Sprintf("Sugar: %fgr", c.sugar)
}

func (c Cake) getName() string {
	return fmt.Sprintf("Name: %s", c.name)
}
func (c Candle) getName() string {
	return fmt.Sprintf("Name: %s", c.name)
}

func main() {

	var tort Cake
	tort.weight=0.6
	tort.sugar = 12.4
	tort.name = "Bananli tort"


	fmt.Println(tort.getName())
	fmt.Println(tort.getSugar())

	var chocolate Candle
	chocolate.weight=0.004
	chocolate.sugar = 25.6
	chocolate.name = "Chocolate with nuts"

	fmt.Println(chocolate.getName())
	fmt.Println(tort.getSugar())
}
