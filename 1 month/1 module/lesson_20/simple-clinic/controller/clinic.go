package controller

import (
	"log"
	"simple-clinic/config"
	"simple-clinic/models"
)

func (c *Conn) CreateClinic(req *models.CreateClinic) (*models.Clinic, error) {

	log.Println(config.Info, "Create clinic resp >>>>> ", req)
	resp, err := c.storage.Clinic().Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) GetByIDClinic(req *models.ClinicPrimaryKey) (*models.Clinic, error) {

	log.Println(config.Info, "GetByID clinic resp >>>>> ", req)
	resp, err := c.storage.Clinic().GetByID(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) GetListClinic(req *models.GetListClinicRequest) (*models.GetListClinicResponse, error) {

	log.Println(config.Info, "GetList clinic resp >>>>> ", req)
	resp, err := c.storage.Clinic().GetList(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) UpdateClinic(req *models.UpdateClinic) (*models.Clinic, error) {

	log.Println(config.Info, "Update clinic resp >>>>> ", req)
	resp, err := c.storage.Clinic().Update(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) DeleteClinic(req *models.ClinicPrimaryKey) error {

	log.Println(config.Info, "Delete clinic resp >>>>> ", req)
	err := c.storage.Clinic().Delete(req)
	if err != nil {
		return err
	}

	return nil
}
