package controller

import (
	"app/models"
	"log"
)

func (c *Connect) CreateCategory(req *models.CreateCategory) (*models.Category, error) {

	log.Println("Create category resp >>>>> ", req)
	resp, err := c.storage.Category().Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (c *Connect) UpdateCategory(req *models.UpadetCategory) (*models.Category, error) {

	log.Println("Update category resp >>>>> ", req)
	resp, err := c.storage.Category().Update(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Connect) GetByIdCategory(req *models.PrimeryKeyCategory) (*models.Category, error) {

	log.Println("GetById category resp >>>>> ", req)
	resp, err := c.storage.Category().GetByID(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Connect) GetListCategory(req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error) {

	log.Println("GetList category resp >>>>> ", req)
	resp, err := c.storage.Category().GetList(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Connect) DeleteCategory(req *models.PrimeryKeyCategory) error {

	log.Println("GetList category resp >>>>> ", req)
	err := c.storage.Category().Delete(req)
	if err != nil {
		return err
	}

	return nil
}
