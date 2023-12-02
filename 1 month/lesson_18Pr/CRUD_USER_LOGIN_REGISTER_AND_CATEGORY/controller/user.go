package controller

import (
	"app/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

type UserController struct {
	Users []models.User
}

func (u *UserController) Login(req models.LoginRequest) models.LoginRespons {
	for _, el := range u.Users {
		fmt.Println("ELEMENT", req.UserName, el.UserName)
		if el.UserName == req.UserName {
			return models.LoginRespons{Message: "Access"}
		}

	}
	return models.LoginRespons{Message: "You are not registered"}
}

func (u *UserController) Register(req models.RegisterRequest) models.RegisterRespons {
	for _, el := range u.Users {
		if el.UserName == req.UserName {
			return models.RegisterRespons{Message: fmt.Sprintf("Bu %s username bazada bor boshqa username kiriting", req.UserName)}
		}
	}
	u.ReadAll("JsonDB/users.json")
	u.Users = append(u.Users, u.CreateUser(models.UserCreate{
		Name:        req.Name,
		SurName:     req.SurName,
		UserName:    req.UserName,
		Password:    req.Password,
		Age:         req.Age,
		DateOfBirth: req.DateOfBirth,
	}))
	u.WriteFile("JsonDB/users.json")
	return models.RegisterRespons{Message: "Access"}
}

func (u *UserController) WriteFile(path string) {
	respons, err := json.MarshalIndent(u.Users, " ", "   ")
	// resp,err := io.ReadAll()
	if err != nil {
		log.Println(err)
	}
	// fmt.Printf(" %s\n", string(respons))

	err = os.WriteFile(path, respons, 0644)
	if err != nil {
		log.Println(err)
	}

}
func (u *UserController) ReadAll(path string) {
	resp, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(resp, &u.Users)
	// resp,err := io.ReadAll()
	if err != nil {
		log.Println(err)
	}
}

/* ***********************************************New User Controller Function*********************************************** */

func NewUserController(users []models.User) UserController {
	return UserController{
		Users: users,
	}
}

/* ***********************************************New User Controller Function*********************************************** */
/* ***********************************************Generate User Method*********************************************** */
func (u *UserController) GenerateUser(limit int) *UserController {
	for i := 0; i < limit; i++ {
		AgeVsDate := models.DateAndAge{}
		AgeVsDate.AgeAndDate(faker.Date())
		u.CreateUser(models.UserCreate{
			Name:        faker.FirstName(),
			SurName:     faker.LastName(),
			UserName:    faker.Username(),
			Password:    faker.Password(),
			Age:         AgeVsDate.Age,
			DateOfBirth: AgeVsDate.Date,
		})
	}
	return u
}

/* ***********************************************Generate User Method*********************************************** */
/* ***********************************************Create User Method*********************************************** */
func (u *UserController) CreateUser(req models.UserCreate) models.User {
	newUser := models.User{
		Id:          uuid.New(),
		Name:        req.Name,
		Surname:     req.SurName,
		UserName:    req.UserName,
		Password:    req.Password,
		Age:         req.Age,
		DateOfBirth: req.DateOfBirth,
		CreatedAt:   time.Now().Format("2006-01-01"),
		UpdatedAt:   time.Now().Format("2006-01-01"),
	}
	u.ReadAll("JsonDB/users.json")
	u.Users = append(u.Users, newUser)
	u.WriteFile("JsonDB/users.json")
	return newUser
}

/* ***********************************************Create User Method*********************************************** */
/* ***********************************************User Get List Method*********************************************** */
func (u *UserController) UserGetList(req models.UserGetListRequest) models.UserGetListRespons {
	fmt.Println(" ********************************************UserGetList Started******************************************** ")
	newLimitUsers := []models.User{}
	if req.Limit == 0 {
		req.Limit = 10
	}
	if req.OffSet+req.Limit > len(u.Users) {
		fmt.Println("Counnt ", len(u.Users))
		for i, el := range u.Users {
			fmt.Println(i+1, "  ", el)
		}

		fmt.Println(" ********************************************UserGetList Finished******************************************** ")
		return models.UserGetListRespons{
			Count: len(u.Users),
			Users: u.Users,
		}
	}
	fmt.Println("Counnt ", len(u.Users))
	for i := req.OffSet; i < req.Limit+req.OffSet; i++ {
		fmt.Println(i+1, u.Users[i])
	}
	newLimitUsers = append(newLimitUsers, u.Users...)
	fmt.Println(" ********************************************UserGetList Finished******************************************** ")
	return models.UserGetListRespons{
		Count: len(u.Users),
		Users: newLimitUsers,
	}
}

/* ***********************************************User Get List Method*********************************************** */
/* ***********************************************User Get By Id Method*********************************************** */
func (u *UserController) UserGetById(req models.UserPrimaryKey) models.User {
	u.ReadAll("JsonDB/users.json")
	// fmt.Println("eeeeeeeeeeeeeee", u.Users)
	for _, el := range u.Users {
		if el.Id == req.Id {
			return el
		}
	}
	return models.User{}
}

/* ***********************************************User Get By Id Method*********************************************** */
/* ***********************************************User Update Method*********************************************** */
func (u *UserController) UserUpdate(req models.UserUpdateRequest) models.User {
	AgeVsDate := models.DateAndAge{}
	AgeVsDate.AgeAndDate(faker.Date())
	for i, el := range u.Users {
		if el.Id == req.Id {
			u.Users[i].Name = req.Name
			u.Users[i].Surname = req.Surname
			u.Users[i].Age = AgeVsDate.Age
			u.Users[i].UserName = req.Username
			u.Users[i].Password = req.Password
			u.Users[i].DateOfBirth = AgeVsDate.Date
			u.Users[i].UpdatedAt = time.Now().Format("2006-01-01")
			u.WriteFile("JsonDB/users.json")
			return u.Users[i]
		}
	}
	return models.User{}
}

/* ***********************************************User Update Method*********************************************** */
/* ***********************************************User Delete Method*********************************************** */
func (u *UserController) UserDelete(req models.UserPrimaryKey) (string, models.User) {

	for index, el := range u.Users {
		if el.Id == req.Id {
			del_user := u.Users[index]
			u.Users = append(u.Users[:index], u.Users[index+1:]...)
			u.WriteFile("JsonDB/users.json")
			return "DELETED", del_user
		}
	}
	return fmt.Sprintf("%s shu ID ega User Topilmadi", req.Id.String()), models.User{}
}

/* ***********************************************User Delete Method*********************************************** */
