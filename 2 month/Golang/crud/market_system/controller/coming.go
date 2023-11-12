package controller

import (
	"errors"
	"log"

	"github.com/asadbekGo/market_system/config"
	"github.com/asadbekGo/market_system/models"
)

func (c *Conn) CreateComing(req *models.CreateComing) (*models.Coming, error) {

	log.Println(config.Info, "Create Coming resp >>>>> ", req)
	resp, err := c.storage.Coming().Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) GetByIDComing(req *models.ComingPrimaryKey) (*models.Coming, error) {

	log.Println(config.Info, "GetByID Coming resp >>>>> ", req)
	resp, err := c.storage.Coming().GetByID(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) GetListComing(req *models.GetListComingRequest) (*models.GetListComingResponse, error) {

	log.Println(config.Info, "GetList Coming resp >>>>> ", req)
	resp, err := c.storage.Coming().GetList(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) UpdateComing(req *models.UpdateComing) (*models.Coming, error) {

	log.Println(config.Info, "Update Coming resp >>>>> ", req)
	rowsAffected, err := c.storage.Coming().Update(req)
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("no rows affected")
	}

	resp, err := c.storage.Coming().GetByID(&models.ComingPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) DeleteComing(req *models.ComingPrimaryKey) error {

	log.Println(config.Info, "Delete Coming resp >>>>> ", req)
	err := c.storage.Coming().Delete(req)
	if err != nil {
		return err
	}

	return nil
}
