package controller

import (
	"app/model"
	"fmt"
	"math/rand"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

type UserController struct {
	Users []model.User
}

func NewUserController(users []model.User) UserController {
	return UserController{
		Users: users,
	}
}

func (u UserController) GetList(req model.GetListUserRequest) model.GetListUserRespons {
	if req.Limit == 0 {
		req.Limit = 10
	}
	var respons = model.GetListUserRespons{}
	if req.Offset+req.Limit > len(u.Users) {
		return model.GetListUserRespons{}
	}

	for i := req.Offset; i < req.Limit+req.Offset; i++ {
		respons.Users = append(respons.Users, u.Users[i])
	}

	respons.Count = len(u.Users)

	return respons

}

func (u UserController) Print(respons model.GetListUserRespons) {
	fmt.Println("Count", respons.Count)
	for index, user := range respons.Users {
		fmt.Println(index+1, user)
	}
}

func UserGenerateData(limit int) []model.User {
	var users []model.User
	for i := 0; i < limit; i++ {
		users = append(users, model.User{
			Id:        uuid.New().String(),
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Age:       rand.Intn(80) + 20,
		})
	}

	return users
}
