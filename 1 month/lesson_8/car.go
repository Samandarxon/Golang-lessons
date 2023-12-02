package main

import "fmt"

type Car struct {
	name  string
	year  int
	price int
}

func generateCarData(limit int) []Car {

	var (
		cars []Car
		year = 2000
	)

	for i := 0; i < limit; i++ {

		year++
		if year > 2024 {
			year = 2000
		}

		cars = append(cars, Car{
			name:  fmt.Sprintf("BMW-%d", i+1),
			year:  year,
			price: 1_000_000,
		})
	}

	return cars
}

func main() {

	cars := generateCarData(100)

	for _, car := range cars {
		fmt.Println(car)
	}

	// carDateFilter(2000, 2005)
}
