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
func carDateFilter(x,y int)[]Car{
	slice:=[]Car{}
	cars := generateCarData(100)
	for _,car :=range cars{
		fmt.Println(car)
	}
	for _,car:=range cars{
		if car.year>=x && car.year<=y{
			// fmt.Println(car)
			slice=append(slice,car)
		}
	}
	return slice
}
func main() {

	// cars := generateCarData(100)

	// for _, car := range cars {
	// 	fmt.Println(car)
	// }

	fmt.Println(carDateFilter(2000, 2005))
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// fmt.Println(carDateFilter(2000, 2000))
	// fmt.Println()
	// fmt.Println(carDateFilter(2006, 2005))
	// fmt.Println()
	// fmt.Println()

}
