package jsonDB

import (
	"app/config"
	"app/models"
	"app/pkg/helpers"
	"app/storage"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cast"
)

type userRepo struct {
	userFile string
}

func NewUserRepo(cfg *config.Config) storage.UserRepoI {

	return &userRepo{userFile: cfg.FilePath + cfg.File.UserFile}
}

func (u *userRepo) Create(req *models.CreateUser) (*models.User, error) {

	data, err := helpers.Read(u.userFile)
	if err != nil {
		return nil, err
	}

	var resp = &models.User{
		Id:        uuid.New().String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Balance:   req.Balance,
	}

	data = append(data, resp)

	err = helpers.Write(u.userFile, data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *userRepo) GetByID(req *models.PrimeryKeyUser) (*models.User, error) {

	data, err := helpers.Read(c.userFile)
	if err != nil {
		return nil, err
	}

	for _, userObj := range data {
		var user = cast.ToStringMap(userObj)
		if cast.ToString(user["id"]) == req.Id {

			return &models.User{
				Id:        cast.ToString(user["id"]),
				FirstName: cast.ToString(user["first_name"]),
				LastName:  cast.ToString(user["last_name"]),
				Balance:   cast.ToInt(user["balance"]),
			}, nil
		}
	}

	return nil, errors.New("data is no rows")
}

func (c *userRepo) GetList(req *models.GetListUserRequest) (*models.GetListUserResponse, error) {
	data, err := helpers.Read(c.userFile)
	if err != nil {
		return nil, err
	}
	// fmt.Println(data)
	if req.Limit == 0 {
		req.Limit = 10
	}

	var resp = &models.GetListUserResponse{}

	for i := req.Offset; i < req.Limit+req.Offset; i++ {
		if len(data) > i {
			var user = cast.ToStringMap(data[i])
			fmt.Println(i+1, user)
			resp.Users = append(resp.Users, &models.User{
				Id:        cast.ToString(user["id"]),
				FirstName: cast.ToString(user["first_name"]),
				LastName:  cast.ToString(user["last_name"]),
				Balance:   cast.ToInt(user["balance"]),
			})
		} else {
			break
		}
	}
	resp.Count = len(data)

	return resp, nil
}

func (c *userRepo) Update(req *models.UpadetUser) (*models.User, error) {

	data, err := helpers.Read(c.userFile)
	if err != nil {
		return nil, err
	}

	for _, userObj := range data {
		var user = cast.ToStringMap(userObj)

		if cast.ToString(user["id"]) == req.Id {
			user["id"] = req.Id
			user["first_name"] = req.FirstName
			user["last_name"] = req.LastName
			user["balance"] = req.Balance

			// data[ind] = user
			break
		}
	}

	err = helpers.Write(c.userFile, data)
	if err != nil {
		return nil, err
	}

	return c.GetByID(&models.PrimeryKeyUser{Id: req.Id})
}

func (c *userRepo) Delete(req *models.PrimeryKeyUser) error {

	data, err := helpers.Read(c.userFile)
	if err != nil {
		return err
	}

	for ind, userObj := range data {
		var user = cast.ToStringMap(userObj)
		if cast.ToString(user["id"]) == req.Id {
			data = append(data[:ind], data[ind+1:]...)
			break
		}
	}

	err = helpers.Write(c.userFile, data)
	if err != nil {
		return err
	}

	return nil
}
