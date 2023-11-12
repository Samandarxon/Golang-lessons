package controller

import (
	"database/sql"
	"log"

	"example.com/m/model"
)

func (c *Controller) ProductCreate(req model.ProductCreate) (*model.Product, error) {
	log.Println("REQ>>>>>>>>>", req)

	resp, err := c.strg.Product().Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Controller) ProductUpdate(req model.ProductUpdate) (*model.Product, error) {
	log.Println("REQ>>>>>>>>>", req)

	resp, err := c.strg.Product().Update(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Controller) ProductDelete(req model.ProductPrimaryKey) (*sql.Result, error) {
	log.Println("REQ>>>>>>>>>", req)

	resp, err := c.strg.Product().Delete(req)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
func (c *Controller) ProductGetList() ([]model.Product, error) {
	// log.Println("REQ>>>>>>>>>", req)

	resp, err := c.strg.Product().GetList()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (c *Controller) ProductGetId(req model.ProductPrimaryKey) (*model.Product, error) {
	// log.Println("REQ>>>>>>>>>", req)

	resp, err := c.strg.Product().GetId(model.ProductPrimaryKey{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}
