package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsonDB"
	"fmt"
	"log"
	"math/rand"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

func main() {
	var cfg = config.ConnectionDataBases()

	jsondb, err := jsonDB.NewJsonDbConnection(&cfg)
	if err != nil {
		panic(err)
	}

	connect := controller.NewController(&cfg, jsondb)

	fmt.Println("Salom", cfg)
	User(connect)
	Category(connect)
	Order(connect)
	Product(connect)

}

func Product(connect *controller.Connect) {
	// Create
	product, err := connect.CreateProduct(&models.CreateProduct{
		Name:         faker.FirstName(),
		Price:        rand.Intn(50_000) + 10_000,
		Discount:     rand.Intn(100) + 10,
		DiscountType: faker.Century(),
		CategoryID:   "d30a4cbc-68f2-4b30-8d54-8a3bc0f7e581",
	})
	if err != nil {
		log.Println("USER", err)
	}
	fmt.Println(product)
	// Update
	// product, err := connect.UpdateProduct(&models.UpadetProduct{
	// 	Id:        "25d812bb-4c8b-4aef-852e-efa8c475789a",
	// 	Name: faker.FirstName(),
	// })
	// if err != nil {
	// 	log.Println("USER", err)
	// }
	// fmt.Println(product)
	// GetList
	// product, err := connect.GetListProduct(&models.GetListProductRequest{
	// 	Offset: 0,
	// })
	// if err != nil {
	// 	log.Println("USER", err)
	// }
	// fmt.Println(product)
	// GetById
	// product, err := connect.GetByIdProduct(&models.PrimeryKeyProduct{
	// 	Id: "25d812bb-4c8b-4aef-852e-efa8c475789a",
	// })
	// if err != nil {
	// 	log.Println("USER", err)
	// }
	// fmt.Println(product)
	// Delete
	// connect.DeleteProduct(&models.PrimeryKeyProduct{
	// 	Id: "050eabe4-87b7-4e10-8610-e656656205a3",
	// })
}

func Order(connect *controller.Connect) {
	// Create
	order, err := connect.CreateOrder(&models.CreateOrder{
		UserId:    "25d812bb-4c8b-4aef-852e-efa8c475789a",
		Sum:       200000,
		SumCount:  20,
		DateTime:  faker.Date(),
		Status:    "Active",
		ProductId: uuid.New().String(),
	})
	if err != nil {
		log.Println("USER", err)
	}
	fmt.Println(order)
	// Update
	// order, err := connect.UpdateOrder(&models.UpadetOrder{
	// 	Id:        "25d812bb-4c8b-4aef-852e-efa8c475789a",
	// 	Name: faker.FirstName(),
	// })
	// if err != nil {
	// 	log.Println("USER", err)
	// }
	// fmt.Println(order)
	// GetList
	// order, err := connect.GetListOrder(&models.GetListOrderRequest{
	// 	Offset: 0,
	// })
	// if err != nil {
	// 	log.Println("USER", err)
	// }
	// fmt.Println(order)
	// GetById
	// order, err := connect.GetByIdOrder(&models.PrimeryKeyOrder{
	// 	Id: "25d812bb-4c8b-4aef-852e-efa8c475789a",
	// })
	// if err != nil {
	// 	log.Println("USER", err)
	// }
	// fmt.Println(order)
	// Delete
	// connect.DeleteOrder(&models.PrimeryKeyOrder{
	// 	Id: "050eabe4-87b7-4e10-8610-e656656205a3",
	// })
}

func Category(connect *controller.Connect) {
	// Create
	category, err := connect.CreateCategory(&models.CreateCategory{
		Name: faker.FirstName(),
	})
	if err != nil {
		log.Println("USER", err)
	}
	fmt.Println(category)
	// Update
	// category, err := connect.UpdateCategory(&models.UpadetCategory{
	// 	Id:        "25d812bb-4c8b-4aef-852e-efa8c475789a",
	// 	Name: faker.FirstName(),
	// })
	// if err != nil {
	// 	log.Println("USER", err)
	// }
	// fmt.Println(category)
	// GetList
	// category, err := connect.GetListCategory(&models.GetListCategoryRequest{
	// 	Offset: 0,
	// })
	// if err != nil {
	// 	log.Println("USER", err)
	// }
	// fmt.Println(category)
	// GetById
	// category, err := connect.GetByIdCategory(&models.PrimeryKeyCategory{
	// 	Id: "25d812bb-4c8b-4aef-852e-efa8c475789a",
	// })
	// if err != nil {
	// 	log.Println("USER", err)
	// }
	// fmt.Println(category)
	// Delete
	// connect.DeleteCategory(&models.PrimeryKeyCategory{
	// 	Id: "050eabe4-87b7-4e10-8610-e656656205a3",
	// })
}
func User(connect *controller.Connect) {
	// Create
	user, err := connect.CreateUser(&models.CreateUser{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Balance:   rand.Intn(40_000) + 10_000,
	})
	if err != nil {
		log.Println("USER", err)
	}
	fmt.Println(user)
	// Update
	// user, err := connect.UpdateUser(&models.UpadetUser{
	// 	Id:        "25d812bb-4c8b-4aef-852e-efa8c475789a",
	// 	FirstName: faker.FirstName(),
	// 	LastName:  faker.LastName(),
	// 	Balance:   rand.Intn(40_000) + 10_000,
	// })
	// if err != nil {
	// 	log.Println("USER", err)
	// }
	// fmt.Println(user)
	// GetList
	// user, err := connect.GetListUser(&models.GetListUserRequest{
	// 	Offset: 0,
	// })
	// if err != nil {
	// 	log.Println("USER", err)
	// }
	// fmt.Println(user)
	// GetById
	// user, err := connect.GetByIdUser(&models.PrimeryKeyUser{
	// 	Id: "25d812bb-4c8b-4aef-852e-efa8c475789a",
	// })
	// if err != nil {
	// 	log.Println("USER", err)
	// }
	// fmt.Println(user)
	// Delete
	// connect.DeleteUser(&models.PrimeryKeyUser{
	// 	Id: "050eabe4-87b7-4e10-8610-e656656205a3",
	// })
}
