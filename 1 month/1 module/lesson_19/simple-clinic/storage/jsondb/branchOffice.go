package jsondb

import (
	"encoding/json"
	"fmt"
	"os"

	"simple-clinic/config"
	"simple-clinic/models"
	"simple-clinic/storage"

	"github.com/bxcodec/faker/v3"

	"github.com/google/uuid"
)

type branchOfficeRepo struct {
	cfg  *config.Config
	file *os.File
}

func NewBranchOfficeRepo(cfg *config.Config, file *os.File) storage.BranchOfficeRepoI {
	return &branchOfficeRepo{cfg: cfg, file: file}
}
func (b *branchOfficeRepo) WriteFile(branch_office []*models.BranchOffice) ([]*models.BranchOffice, error) {

	body, err := json.MarshalIndent(branch_office, "", "  ")
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(b.cfg.BranchOfficePath, body, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return branch_office, nil
}

func (b *branchOfficeRepo) ReadAll() ([]*models.BranchOffice, error) {

	body, err := os.ReadFile(b.cfg.BranchOfficePath)
	if err != nil {
		return nil, err
	}
	var branch_office []*models.BranchOffice

	err = json.Unmarshal(body, &branch_office)
	if err != nil {
		return nil, err
	}

	return branch_office, nil
}
func (b *branchOfficeRepo) Create(req *models.CreateBranchOffice) (*models.BranchOffice, error) {

	branch_office, err := b.ReadAll()
	if err != nil {
		return nil, err
	}

	var resp = &models.BranchOffice{
		BranchOfficeID: uuid.New().String(),
		ClinicID:       req.ClinicID,
		Addres:         "Termiz",
		Phone:          faker.Phonenumber(),
	}

	branch_office = append(branch_office, resp)
	_, err = b.WriteFile(branch_office)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *branchOfficeRepo) Update(req *models.UpdateBranchOffice) (*models.BranchOffice, error) {

	branch_office, err := b.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, el := range branch_office {
		if el.BranchOfficeID == req.BranchOfficeID {
			el.ClinicID = req.ClinicID
			el.Addres = req.Addres
			el.Phone = faker.Phonenumber()
			_, err = b.WriteFile(branch_office)
			if err != nil {
				return nil, err
			}
			return el, nil
		}

	}
	return &models.BranchOffice{}, nil
}

func (b *branchOfficeRepo) GetByID(req *models.PrimeryKeyBranchOfficeID) (*models.BranchOffice, error) {

	branch_office, err := b.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, el := range branch_office {
		if el.BranchOfficeID == req.BranchOfficeID {
			return el, nil
		}

	}
	return &models.BranchOffice{}, nil
}

func (b *branchOfficeRepo) Delete(req *models.PrimeryKeyBranchOfficeID) (string, error) {

	branch_office, err := b.ReadAll()
	if err != nil {
		return "Bazada xatolik yuz berdi", err
	}

	for index, el := range branch_office {
		if el.BranchOfficeID == req.BranchOfficeID {
			branch_office = append(branch_office[:index], branch_office[index+1:]...)
			b.WriteFile(branch_office)
			return "DELETED.....", nil
		}

	}
	return fmt.Sprintf("%s Shu IDga ega clinic topilmadi...", req.BranchOfficeID), nil
}
