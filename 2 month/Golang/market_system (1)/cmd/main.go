package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"example.com/m/config"
	"example.com/m/controller"
	"example.com/m/model"
	"example.com/m/storage/postgres"
)

func main() {

	cfg := config.Load()

	strg, err := postgres.NewPostgressConnection(&cfg)
	if err != nil {
		panic(err)
	}

	conn := controller.NewController(&cfg, strg)
	Create(conn)
	// Update(conn)
	// Delete(conn, "8c4fc788-f9a8-4563-b5e6-4e6db0480bc7")
	// GetList(conn)
	// GetId(conn, "8ba2a6be-171e-4ec7-8783-8b382799ab3e")
	defer conn.CloseDb()

}

func Create(conn *controller.Controller) {
	product, err := conn.ProductCreate(model.ProductCreate{
		Name:      "Iphone",
		Price:     10000,
		Barcode:   "12335465",
		Image_url: "123456",
	})
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(product)

}
func GetList(conn *controller.Controller) {
	product, err := conn.ProductGetList()
	if err != nil {
		log.Println(err.Error())
		return
	}
	for _, v := range product {
		fmt.Printf("%+v\n", v)
	}

}
func Update(conn *controller.Controller) {
	product, err := conn.ProductUpdate(model.ProductUpdate{
		Id:        "10e4822d-ee68-40ab-b6db-5f61310ec1da",
		Name:      "Iphone 14 Pro Max",
		Price:     14000,
		Barcode:   "1233657865",
		Image_url: "87979879",
	})
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(product)

}

func Delete(conn *controller.Controller, id string) {
	product, err := conn.ProductDelete(model.ProductPrimaryKey{
		Id: id,
	})
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(product)

}

func GetId(conn *controller.Controller, id string) {
	product, err := conn.ProductGetId(model.ProductPrimaryKey{
		Id: id,
	})
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(product)

}
