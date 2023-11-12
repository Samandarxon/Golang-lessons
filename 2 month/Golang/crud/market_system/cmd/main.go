package main

import (
	"fmt"
	"log"
	"math/rand"

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
	// product(con)
	// remaining(con)
	// branch(con)
	coming(con)
}
func remaining(con *controller.Conn) {

	for i := 0; i < 100; i++ {
		con.CreateRemaining(&models.CreateRemaining{
			Branch_id:  "81ae69f4-5323-4314-98de-6baa09afd276",
			Product_id: "4614b0e4-1997-40b5-84c1-c2661edf76d2",
			Quantity:   rand.Intn(20) + 2,
			Price:      float64(rand.Intn(312) + 10000),
		})
	}

	// resp, err := con.GetListRemaining(&models.GetListRemainingRequest{
	// 	Offset: 0,
	// 	Limit:  100,
	// 	Search: "",
	// })

	// if err != nil {
	// 	log.Println("Error while GetListRemaining >>> " + err.Error())
	// 	return
	// }

	// fmt.Println("Remaining Count:", resp.Count)
	// for _, remaining := range resp.Remainings {
	// 	fmt.Println(strings.Repeat("-", 100))
	// 	fmt.Println(remaining.Id)
	// 	fmt.Println(remaining.Branch_id)
	// 	fmt.Println(remaining.Product_id)
	// 	fmt.Println(remaining.Price)
	// 	fmt.Println(remaining.Quantity)
	// 	fmt.Println(remaining.CreatedAt)
	// 	fmt.Println(remaining.UpdatedAt)
	// 	fmt.Println(strings.Repeat("-", 100))

	// }

	// resp, err := con.UpdateRemaining(&models.UpdateRemaining{
	// 	Id:         "22ba6800-612a-42fc-9890-b1e1b322a193",
	// 	Branch_id:  "a485bd50-c087-470a-be38-7b689acabb30",
	// 	Product_id: "",
	// 	Price:      200,
	// })

	// if err != nil {
	// 	log.Println("Error while UpdateCategory >>> " + err.Error())
	// 	return
	// }
	// resp, err := con.GetByIDRemaining(&models.RemainingPrimaryKey{
	// 	Id: "22ba6800-612a-42fc-9890-b1e1b322a193",
	// })

	// if err != nil {
	// 	log.Println("Error while GetByIDRemaining >>> " + err.Error())
	// 	return
	// }

	// fmt.Println(resp)
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

func branch(con *controller.Conn) {

	// for i := 0; i < 100; i++ {
	// 	res, err := con.CreateBranch(&models.CreateBranch{
	// 		Name:    faker.FirstName(),
	// 		Address: faker.DomainName(),
	// 	})
	// 	fmt.Println(res, err)
	// }

	// resp, err := con.GetListBranch(&models.GetListBranchRequest{
	// 	Offset: 0,
	// 	Limit:  100,
	// 	Search: "m",
	// })

	// if err != nil {
	// 	log.Println("Error while GetListProduct >>> " + err.Error())
	// 	return
	// }

	// fmt.Println("Product Count:", resp.Count)
	// for _, product := range resp.Branchs {
	// 	fmt.Println(product.Name)
	// }

	// resp, err := con.UpdateBranch(&models.UpdateBranch{
	// 	Id:      "f85a22f9-fd29-48dd-835d-fa8d1573142c",
	// 	Name:    "samandarxon",
	// 	Address: "O'zbekiston",
	// })

	resp, err := con.GetByIDBranch(&models.BranchPrimaryKey{
		Id: "f85a22f9-fd29-48dd-835d-fa8d1573142c",
	})

	if err != nil {
		log.Println("Error while UpdateCategory >>> " + err.Error())
		return
	}

	fmt.Println(resp)
}

func coming(con *controller.Conn) {

	// res, err := con.CreateComing(&models.CreateComing{
	// 	Remaining_Id: "c3b217d7-875b-4de1-8f43-75ca35e724bd",
	// })
	// fmt.Println(res, err)

	// resp, err := con.GetListComing(&models.GetListComingRequest{
	// 	Offset: 0,
	// 	Limit:  100,
	// 	Search: "",
	// })

	// if err != nil {
	// 	log.Println("Error while GetListProduct >>> " + err.Error())
	// 	return
	// }

	// fmt.Println("Product Count:", resp.Count)
	// for _, product := range resp.Comings {
	// 	fmt.Println(product)
	// }

	// resp, err := con.UpdateComing(&models.UpdateComing{
	// 	Id:           "536b9e56-7dd0-4af0-8d27-f58ab3b05ea6",
	// 	Remaining_Id: "bec89256-4039-4f40-acb7-e745da079ec3",
	// })

	resp, err := con.GetByIDComing(&models.ComingPrimaryKey{
		Id: "536b9e56-7dd0-4af0-8d27-f58ab3b05ea6",
	})

	if err != nil {
		log.Println("Error while UpdateCategory >>> " + err.Error())
		return
	}

	fmt.Println(resp)
}
