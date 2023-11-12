package controller

import (
	"example.com/m/config"
	"example.com/m/storage"
)

type Controller struct {
	cfg  *config.Config
	strg storage.StorageI
}

func NewController(cfg *config.Config, strg storage.StorageI) *Controller {
	return &Controller{
		cfg:  cfg,
		strg: strg,
	}
}

func (c *Controller) CloseDb() error {
	err := c.strg.CloseDB()
	if err != nil {
		return err
	}
	return nil
}
