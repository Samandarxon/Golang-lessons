package main

import (
	"app/controller"
	"app/model"
)

func main() {
	users := controller.UserGenerateData(1000)
	userController := controller.NewUserController(users)
	userController.Print(userController.GetList(model.GetListUserRequest{
		Offset: 0,
		Limit:  20,
	}))
	respons := userController.GetList(model.GetListUserRequest{
		Offset: 0,
		Limit:  5,
	})
	userController.Print(respons)

}
