package controller

import (
	"errors"
	"log"

	"github.com/asadbekGo/market_system/config"
	"github.com/asadbekGo/market_system/models"
)

func (c *Conn) CreateProduct(req *models.CreateProduct) (*models.Product, error) {

	log.Println(config.Info, "Create Product resp >>>>> ", req)
	resp, err := c.storage.Product().Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) GetByIDProduct(req *models.ProductPrimaryKey) (*models.Product, error) {

	log.Println(config.Info, "GetByID Product resp >>>>> ", req)
	resp, err := c.storage.Product().GetByID(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) GetListProduct(req *models.GetListProductRequest) (*models.GetListProductResponse, error) {

	log.Println(config.Info, "GetList Product resp >>>>> ", req)
	resp, err := c.storage.Product().GetList(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) UpdateProduct(req *models.UpdateProduct) (*models.Product, error) {

	log.Println(config.Info, "Update Product resp >>>>> ", req)
	rowsAffected, err := c.storage.Product().Update(req)
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("no rows affected")
	}

	resp, err := c.storage.Product().GetByID(&models.ProductPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) DeleteProduct(req *models.ProductPrimaryKey) error {

	log.Println(config.Info, "Delete Product resp >>>>> ", req)
	err := c.storage.Product().Delete(req)
	if err != nil {
		return err
	}

	return nil
}
