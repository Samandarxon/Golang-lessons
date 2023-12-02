package controller

import (
	"log"
	"simple-clinic/config"
	"simple-clinic/models"
)

func (c *Conn) CreateBranch(req *models.CreateBranch) (*models.Branch, error) {

	log.Println(config.Info, "Create branch resp >>>>> ", req)
	resp, err := c.storage.Branch().Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) GetByIDBranch(req *models.BranchPrimaryKey) (*models.Branch, error) {

	log.Println(config.Info, "GetByID branch resp >>>>> ", req)
	resp, err := c.storage.Branch().GetByID(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) GetListBranch(req *models.GetListBranchRequest) (*models.GetListBranchResponse, error) {

	log.Println(config.Info, "GetList branch resp >>>>> ", req)
	resp, err := c.storage.Branch().GetList(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) UpdateBranch(req *models.UpdateBranch) (*models.Branch, error) {

	log.Println(config.Info, "Update branch resp >>>>> ", req)
	resp, err := c.storage.Branch().Update(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) DeleteBranch(req *models.BranchPrimaryKey) error {

	log.Println(config.Info, "Delete branch resp >>>>> ", req)
	err := c.storage.Branch().Delete(req)
	if err != nil {
		return err
	}

	return nil
}
