package main

import (
	"math"
	"fmt"
	
)

type Color struct {
	R, G, B uint8
}

type Car struct {
	model string
	color Color
	fuel int
	speed_km int
}

func (c Color) getColor() int {
	return int(c.R) + int(c.G) + int(c.B)
}

func (c Car) getName() string {
	return "Car Name: " + c.model
}

func (c Car) distanceDrive() float64{
	distance := math.Ceil(float64(c.fuel)*100/float64(c.speed_km))
	return distance
}

func main() {

	var dodgerblue = Color{
		R: 30, G: 144, B: 255,
	}

	var bmw = Car{
		model: "BMW X6",
		color: dodgerblue,
		fuel:100,
		speed_km:15,
	}

	fmt.Println(bmw.color.getColor())
	fmt.Println(bmw.getName())
	fmt.Println(bmw.distanceDrive())

}
