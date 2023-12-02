package controller

import (
	"app/models"
	"log"
)

func (c *Connect) CreateUser(req *models.CreateUser) (*models.User, error) {

	log.Println("Create user resp >>>>> ", req)
	resp, err := c.storage.User().Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (c *Connect) UpdateUser(req *models.UpadetUser) (*models.User, error) {

	log.Println("Update user resp >>>>> ", req)
	resp, err := c.storage.User().Update(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Connect) GetByIdUser(req *models.PrimeryKeyUser) (*models.User, error) {

	log.Println("GetById user resp >>>>> ", req)
	resp, err := c.storage.User().GetByID(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Connect) GetListUser(req *models.GetListUserRequest) (*models.GetListUserResponse, error) {

	log.Println("GetList user resp >>>>> ", req)
	resp, err := c.storage.User().GetList(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Connect) DeleteUser(req *models.PrimeryKeyUser) error {

	log.Println("GetList user resp >>>>> ", req)
	err := c.storage.User().Delete(req)
	if err != nil {
		return err
	}

	return nil
}
