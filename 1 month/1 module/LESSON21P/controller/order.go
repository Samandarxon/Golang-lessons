package controller

import (
	"app/models"
	"log"
)

func (c *Connect) CreateOrder(req *models.CreateOrder) (*models.Order, error) {

	log.Println("Create order resp >>>>> ", req)
	resp, err := c.storage.Order().Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (c *Connect) UpdateOrder(req *models.UpadetOrder) (*models.Order, error) {

	log.Println("Update order resp >>>>> ", req)
	resp, err := c.storage.Order().Update(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Connect) GetByIdOrder(req *models.PrimeryKeyOrder) (*models.Order, error) {

	log.Println("GetById order resp >>>>> ", req)
	resp, err := c.storage.Order().GetByID(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Connect) GetListOrder(req *models.GetListOrderRequest) (*models.GetListOrderResponse, error) {

	log.Println("GetList order resp >>>>> ", req)
	resp, err := c.storage.Order().GetList(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Connect) DeleteOrder(req *models.PrimeryKeyOrder) error {

	log.Println("GetList order resp >>>>> ", req)
	err := c.storage.Order().Delete(req)
	if err != nil {
		return err
	}

	return nil
}
