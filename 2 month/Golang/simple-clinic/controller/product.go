package controller

import (
	"app/models"
	"log"
)

func (c *Connect) CreateProduct(req *models.CreateProduct) (*models.Product, error) {

	log.Println("Create product resp >>>>> ", req)
	resp, err := c.storage.Product().Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (c *Connect) UpdateProduct(req *models.UpadetProduct) (*models.Product, error) {

	log.Println("Update product resp >>>>> ", req)
	resp, err := c.storage.Product().Update(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Connect) GetByIdProduct(req *models.PrimeryKeyProduct) (*models.Product, error) {

	log.Println("GetById product resp >>>>> ", req)
	resp, err := c.storage.Product().GetByID(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Connect) GetListProduct(req *models.GetListProductRequest) (*models.GetListProductResponse, error) {

	log.Println("GetList product resp >>>>> ", req)
	resp, err := c.storage.Product().GetList(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Connect) DeleteProduct(req *models.PrimeryKeyProduct) error {

	log.Println("GetList product resp >>>>> ", req)
	err := c.storage.Product().Delete(req)
	if err != nil {
		return err
	}

	return nil
}
