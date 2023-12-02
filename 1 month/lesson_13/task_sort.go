package main

import (
	"fmt"
	"sort"
)

const (
	NAME = iota
	AGE
)

type User struct {
	Name string
	Age  int
}

func sorted(users []User, x int) []User{
	switch x {
	case NAME:
		sort.Slice(users, func(i, j int) bool {
			return users[i].Name < users[j].Name
		})
	case AGE:
		sort.Slice(users, func(i, j int) bool {
			return users[i].Age < users[j].Age
		})
		
	}
	return users
}
func main() {

	// var nums = []int{3, 5, 6, 12, 3, 56, 8, 2}

	// sort.Slice(nums, func(x, y int) bool {
	//   return nums[x] < nums[y]
	// })

	// fmt.Println(nums)

	var users = []User{
		{Name: "Yasharbek", Age: 20},
		{Name: "Zafar", Age: 18},
		{Name: "Ilyos", Age: 26},
		{Name: "Akrom", Age: 18},
		{Name: "Ilgor", Age: 16},
		{Name: "Javohir", Age: 16},
		{Name: "Fayzulloh", Age: 21},
		{Name: "Izzat", Age: 22},
		{Name: "Baxodir", Age: 22},
	}

	sorted(users, NAME)
	for _, user := range users {
		fmt.Println(user)
	}
	fmt.Println()
	sorted(users, AGE)

	for _, user := range users {
		fmt.Println(user)
	}

}
