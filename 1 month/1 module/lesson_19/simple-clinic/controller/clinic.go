package controller

import (
	"log"
	"simple-clinic/models"
)

func (c *conn) CreateClinic(req *models.CreateClinic) (*models.Clinic, error) {

	log.Println("Create clinin resp")
	resp, err := c.storage.Clinic().Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (c *conn) UpdateClinic(req *models.UpdateClinic) (*models.Clinic, error) {
	log.Println("Update clinic resp")
	resp, err := c.storage.Clinic().Update(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *conn) GetByIDClinic(req *models.PrimeryKeyClinicID) (*models.Clinic, error) {
	log.Println("Update clinic resp")
	resp, err := c.storage.Clinic().GetByID(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *conn) DeleteClinic(req *models.PrimeryKeyClinicID) (string, error) {
	log.Println("Update clinic resp")
	resp, err := c.storage.Clinic().Delete(req)
	if err != nil {
		return "Bazada xatolik yuz berdi", err
	}
	return resp, nil
}
