package main

import (
	"app/controller"
	"app/models"
	"fmt"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

func main() {

	var Products []models.Product

	productConnection := controller.NewProductController(Products)

	_ = productConnection.CreateProduct(models.ProductCreate{
		Id:       uuid.New(),
		Name:     "Apple",
		Price:    10000,
		Quantity: 10,
	})
	_ = productConnection.CreateProduct(models.ProductCreate{
		Id:       uuid.New(),
		Name:     "Acer",
		Price:    3000,
		Quantity: 12,
	})

	products := productConnection.ProductGetList(models.ProductGetListRequest{})

	productg := productConnection.ProdcutGetById(models.ProductPrimaryKey{
		Id: products.Products[0].Id,
	})

	update := productConnection.ProductUpdate(models.ProductUpdate{
		Id:       productg.Id,
		Price:    2000,
		Quantity: 30,
	})
	deleted := productConnection.ProductDelete(models.ProductPrimaryKey{
		Id: productg.Id,
	})
	// fmt.Println("All", products)
	fmt.Println()
	fmt.Println("GetByID", productg)
	fmt.Println()
	fmt.Println("update", update)
	fmt.Println()
	// fmt.Println("All", products)
	fmt.Println()
	fmt.Println(deleted)
	fmt.Println()
	// fmt.Println("All", products)

	/* ***********************************************HomeWork User Started*********************************************** */

	userController := controller.NewUserController([]models.User{})

	// /* ***********************************************HomeWork User Created*********************************************** */

	_ = *userController.GenerateUser(4)

	// /* ***********************************************HomeWork User Created*********************************************** */

	// /* ***********************************************HomeWork User GetList*********************************************** */

	users := userController.UserGetList(models.UserGetListRequest{
		OffSet: 0,
		Limit:  20,
	})

	// userController.UserGetList(models.UserGetListRequest{
	// 	OffSet: 10,
	// 	Limit:  5,
	// })

	// /* ***********************************************HomeWork User GetList *********************************************** */

	// /* ***********************************************HomeWork User GetById *********************************************** */

	user := userController.UserGetById(models.UserPrimaryKey{Id: users.Users[2].Id})

	fmt.Println("UserGetById", user)

	// /* ***********************************************HomeWork User GetById *********************************************** */

	// /* ***********************************************HomeWork User UserUpdate *********************************************** */

	userUpdate := userController.UserUpdate(models.UserUpdateRequest{
		Id:       user.Id,
		Name:     "Samandarxon",
		Surname:  "Temirov",
		Username: "SAmTEm",
		Password: faker.Password(),
	})
	fmt.Println("User Update", userUpdate)
	userController.UserGetList(models.UserGetListRequest{
		OffSet: 0,
	})

	// /* ***********************************************HomeWork User UserUpdate *********************************************** */
	// /* ***********************************************HomeWork User UserUpdate *********************************************** */

	str, userdel := userController.UserDelete(models.UserPrimaryKey{Id: user.Id})
	fmt.Println("User Delete", str, userdel)
	userController.UserGetList(models.UserGetListRequest{
		OffSet: 0,
	})

	/* ***********************************************HomeWork User UserUpdate *********************************************** */
	/* ***********************************************HomeWork User Finished*********************************************** */
	/********************************************************************************************************************************/
	/********************************************************************************************************************************/
	/********************************************************************************************************************************/
	/********************************************************************************************************************************/
	/********************************************************************************************************************************/
	/* ***********************************************User Login*********************************************** */
	// login := userController.Login(models.LoginRequest{
	// 	UserName: "SAm",
	// 	Password: faker.Password(),
	// })
	// fmt.Println(login.Message)
	// login = userController.Login(models.LoginRequest{
	// 	UserName: users.Users[1].UserName,
	// 	Password: faker.Password(),
	// })
	// fmt.Println(login.Message)

	// /* ***********************************************User Login*********************************************** */
	// /* ***********************************************User Register*********************************************** */
	// AgeVsDate := models.DateAndAge{}
	// AgeVsDate.AgeAndDate(faker.Date())
	// register := userController.Register(models.RegisterRequest{
	// 	Name:        faker.FirstName(),
	// 	SurName:     faker.LastName(),
	// 	UserName:    "asdFRFSDasdas",
	// 	Password:    faker.Password(),
	// 	Age:         AgeVsDate.Age,
	// 	DateOfBirth: AgeVsDate.Date,
	// })
	// fmt.Println("Register", register.Message)

	// AgeVsDate = models.DateAndAge{}
	// AgeVsDate.AgeAndDate(faker.Date())
	// register = userController.Register(models.RegisterRequest{
	// 	Name:        faker.FirstName(),
	// 	SurName:     faker.LastName(),
	// 	UserName:    users.Users[1].UserName,
	// 	Password:    faker.Password(),
	// 	Age:         AgeVsDate.Age,
	// 	DateOfBirth: AgeVsDate.Date,
	// })

	// fmt.Println("Register", register.Message)

	/* ***********************************************User Register*********************************************** */

	/********************************************************************************************************************************/
	/********************************************************************************************************************************/
	/********************************************************************************************************************************/
	/********************************************************************************************************************************/
	/********************************************************************************************************************************/

	/* ***********************************************HomeWork Category Started*********************************************** */

	categoryController := controller.NewCategoryController([]models.Category{})

	/* ***********************************************HomeWork Category Created*********************************************** */

	_ = categoryController.GenerateCategory(5)

	/* ***********************************************HomeWork Category Created*********************************************** */

	/* ***********************************************HomeWork Category GetList*********************************************** */

	// categoryController.CategoryGetList(models.CategoryGetListRequest{
	// 	OffSet: 0,
	// 	Limit:  20,
	// })

	// categoryController.CategoryGetList(models.CategoryGetListRequest{
	// 	OffSet: 10,
	// 	Limit:  5,
	// })

	/* ***********************************************HomeWork Category GetList *********************************************** */

	/* ***********************************************HomeWork Category GetById *********************************************** */
	category := categoryController.CategoryGetById(models.CategoryPrimaryKey{Id: categoryController.Categorys[0].Id})
	fmt.Println("CategoryGetById", category)

	/* ***********************************************HomeWork Category GetById *********************************************** */

	/* ***********************************************HomeWork Category CategoryUpdate *********************************************** */

	categoryUpdate := categoryController.CategoryUpdate(models.CategoryUpdateRequest{
		Id:   category.Id,
		Name: "Samandarxon",
	})
	fmt.Println("Category Update", categoryUpdate)
	// categoryController.CategoryGetList(models.CategoryGetListRequest{
	// 	OffSet: 0,
	// })

	/* ***********************************************HomeWork Category CategoryUpdate *********************************************** */
	/* ***********************************************HomeWork Category CategoryUpdate *********************************************** */

	str, categorydel := categoryController.CategoryDelete(models.CategoryPrimaryKey{Id: category.Id})
	fmt.Println("Category Delete", str, categorydel)
	categoryController.CategoryGetList(models.CategoryGetListRequest{
		OffSet: 0,
		Limit:  20,
	})
	/* ***********************************************HomeWork Category CategoryUpdate *********************************************** */
	/* ***********************************************HomeWork Category Finished*********************************************** */

}
