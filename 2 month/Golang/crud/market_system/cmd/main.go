package main

import (
	"fmt"
	"log"

	"github.com/asadbekGo/market_system/config"
	"github.com/asadbekGo/market_system/controller"
	"github.com/asadbekGo/market_system/models"
	"github.com/asadbekGo/market_system/storage/postgres"
)

func main() {

	var cfg = config.Load()

	pgStorage, err := postgres.NewConnectionPostgres(&cfg)
	if err != nil {
		panic(err)
	}

	con := controller.NewController(&cfg, pgStorage)

	// category(con)
	product(con)
}

func category(con *controller.Conn) {

	// for i := 0; i < 100; i++ {
	// 	respons, err := con.CreateCategory(&models.CreateCategory{
	// 		Title:    faker.FirstName(),
	// 		ParentID: "",
	// 	})
	// 	fmt.Println(respons, err)
	// }

	resp, err := con.GetListCategory(&models.GetListCategoryRequest{
		Offset: 0,
		Limit:  100,
		Search: "m",
	})

	if err != nil {
		log.Println("Error while GetListCategory >>> " + err.Error())
		return
	}

	// fmt.Println("Category Count:", resp.Count)
	// for _, category := range resp.Categories {
	// 	fmt.Println(category.Title)
	// }

	// resp, err := con.UpdateCategory(&models.UpdateCategory{
	// 	Id:       "64c5ebb8-6b43-406e-b285-a3009f9cf3e9",
	// 	Title:    "JUBAJUBA",
	// 	ParentID: "",
	// })

	// if err != nil {
	// 	log.Println("Error while UpdateCategory >>> " + err.Error())
	// 	return
	// }

	fmt.Println(resp)
}

func product(con *controller.Conn) {

	// for i := 0; i < 100; i++ {
	// 	res, err := con.CreateProduct(&models.CreateProduct{
	// 		Name:        faker.FirstName(),
	// 		Price:       cast.ToFloat64(faker.GetPrice()),
	// 		Category_id: "5a01dade-6ac0-4c13-a167-c5e999d81477",
	// 	})
	// 	fmt.Println(res, err)
	// }

	// resp, err := con.GetListProduct(&models.GetListProductRequest{
	// 	Offset: 0,
	// 	Limit:  100,
	// 	Search: "m",
	// })

	// if err != nil {
	// 	log.Println("Error while GetListProduct >>> " + err.Error())
	// 	return
	// }

	// fmt.Println("Product Count:", resp.Count)
	// for _, product := range resp.Products {
	// 	fmt.Println(product.Name)
	// }

	// resp, err := con.UpdateProduct(&models.UpdateProduct{
	// 	Id:          "0778e252-9740-403b-8def-f7b436eaaeb5",
	// 	Name:        "samandarxon",
	// 	Price:       2342342342,
	// 	Category_id: "6f71fdbe-3cd3-4d1a-8d22-cc3f0d031b79",
	// })

	resp, err := con.GetByIDProduct(&models.ProductPrimaryKey{
		Id: "0778e252-9740-403b-8def-f7b436eaaeb5",
	})

	if err != nil {
		log.Println("Error while UpdateCategory >>> " + err.Error())
		return
	}

	fmt.Println(resp)
}
