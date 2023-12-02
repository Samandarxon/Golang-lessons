package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/bxcodec/faker/v3"
)

type User struct {
	Name     string
	Balance  int
	Products []Product
	TotalSum int
}

type Product struct {
	Name  string
	Price int
	Count int
}

func Cassir(user User) {

	var totalSum int

	for _, product := range user.Products {
		totalSum += product.Price * product.Count
	}
}

func main() {

	var users []User

	for i := 0; i < 200; i++ {
		var user = User{
			Name:    faker.FirstName(),
			Balance: rand.Intn(1000) * 100,
		}

		var products []Product

		for i := 0; i < rand.Intn(30)+1; i++ {
			products = append(products, Product{
				Name:  faker.FirstName(),
				Price: rand.Intn(1000) * 10,
				Count: rand.Intn(100),
			})
		}

		user.Products = products

		users = append(users, user)
	}

	userTime := time.Now()

	var wg = sync.WaitGroup{}
	for ind, user := range users {
		go func(wg *sync.WaitGroup, object User, index int) {
			defer wg.Done()
			var totalSum int
			for _, product := range object.Products {
				totalSum += product.Price * product.Count
			}

			users[index].TotalSum = totalSum
		}(&wg, user, ind)
		wg.Add(1)
	}
	wg.Wait()

	for _, user := range users {
		fmt.Println(user.TotalSum)
	}

	fmt.Println("userTime:", time.Since(userTime))
}
