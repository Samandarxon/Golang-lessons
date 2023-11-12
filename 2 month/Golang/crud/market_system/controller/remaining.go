package controller

import (
	"errors"
	"log"

	"github.com/asadbekGo/market_system/config"
	"github.com/asadbekGo/market_system/models"
)

func (c *Conn) CreateRemaining(req *models.CreateRemaining) (*models.Remaining, error) {

	log.Println(config.Info, "Create Remaining resp >>>>> ", req)
	resp, err := c.storage.Remaining().Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) GetByIDRemaining(req *models.RemainingPrimaryKey) (*models.Remaining, error) {

	log.Println(config.Info, "GetByID Remaining resp >>>>> ", req)
	resp, err := c.storage.Remaining().GetByID(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) GetListRemaining(req *models.GetListRemainingRequest) (*models.GetListRemainingResponse, error) {

	log.Println(config.Info, "GetList Remaining resp >>>>> ", req)
	resp, err := c.storage.Remaining().GetList(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) UpdateRemaining(req *models.UpdateRemaining) (*models.Remaining, error) {

	log.Println(config.Info, "Update Remaining resp >>>>> ", req)
	rowsAffected, err := c.storage.Remaining().Update(req)
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("no rows affected")
	}

	resp, err := c.storage.Remaining().GetByID(&models.RemainingPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) DeleteRemaining(req *models.RemainingPrimaryKey) error {

	log.Println(config.Info, "Delete Remaining resp >>>>> ", req)
	err := c.storage.Remaining().Delete(req)
	if err != nil {
		return err
	}

	return nil
}
