package controller

import (
	"log"
	"simple-clinic/models"
)

func (c *conn) CreateBranchOffice(req *models.CreateBranchOffice) (*models.BranchOffice, error) {

	log.Println("Create BranchOffice resp")
	resp, err := c.storage.BranchOffice().Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (c *conn) UpdateBranchOffice(req *models.UpdateBranchOffice) (*models.BranchOffice, error) {
	log.Println("Update BranchOffice resp")
	resp, err := c.storage.BranchOffice().Update(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *conn) GetByIDBranchOffice(req *models.PrimeryKeyBranchOfficeID) (*models.BranchOffice, error) {
	log.Println("Update BranchOffice resp")
	resp, err := c.storage.BranchOffice().GetByID(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *conn) DeleteBranchOffice(req *models.PrimeryKeyBranchOfficeID) (string, error) {
	log.Println("Update BranchOffice resp")
	resp, err := c.storage.BranchOffice().Delete(req)
	if err != nil {
		return "Bazada xatolik yuz berdi", err
	}
	return resp, nil
}
