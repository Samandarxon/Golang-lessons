package controller

import (
	"log"
	"simple-clinic/config"
	"simple-clinic/models"
)

func (c *Conn) CreateCashier(req *models.CreateCashier) (*models.Cashier, error) {

	log.Println(config.Info, "Create clinic resp >>>>> ", req)
	resp, err := c.storage.Cashier().Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) GetByIDCashier(req *models.CashierPrimaryKey) (*models.Cashier, error) {

	log.Println(config.Info, "GetByID clinic resp >>>>> ", req)
	resp, err := c.storage.Cashier().GetByID(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) GetListCashier(req *models.GetListCashierRequest) (*models.GetListCashierResponse, error) {

	log.Println(config.Info, "GetList clinic resp >>>>> ", req)
	resp, err := c.storage.Cashier().GetList(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) UpdateCashier(req *models.UpdateCashier) (*models.Cashier, error) {

	log.Println(config.Info, "Update clinic resp >>>>> ", req)
	resp, err := c.storage.Cashier().Update(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) DeleteCashier(req *models.CashierPrimaryKey) error {

	log.Println(config.Info, "Delete clinic resp >>>>> ", req)
	err := c.storage.Cashier().Delete(req)
	if err != nil {
		return err
	}

	return nil
}
