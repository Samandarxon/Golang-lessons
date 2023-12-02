package jsondb

import (
	"encoding/json"
	"fmt"
	"os"

	"simple-clinic/config"
	"simple-clinic/models"
	"simple-clinic/storage"

	"github.com/google/uuid"
)

type clinicRepo struct {
	cfg  *config.Config
	file *os.File
}

func NewClinicRepo(cfg *config.Config, file *os.File) storage.ClinicRepoI {
	return &clinicRepo{cfg: cfg, file: file}
}
func (c *clinicRepo) WriteFile(clinics []*models.Clinic) ([]*models.Clinic, error) {

	body, err := json.MarshalIndent(clinics, "", "  ")
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(c.cfg.ClinicPath, body, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return clinics, nil
}

func (c *clinicRepo) ReadAll() ([]*models.Clinic, error) {

	body, err := os.ReadFile(c.cfg.ClinicPath)
	if err != nil {
		return nil, err
	}
	var clinics []*models.Clinic

	err = json.Unmarshal(body, &clinics)
	if err != nil {
		return nil, err
	}

	return clinics, nil
}
func (c *clinicRepo) Create(req *models.CreateClinic) (*models.Clinic, error) {

	clinics, err := c.ReadAll()
	if err != nil {
		return nil, err
	}

	var resp = &models.Clinic{
		ClinicID: uuid.New().String(),
		Name:     req.Name,
		Logo:     req.Logo,
		City:     req.City,
	}

	clinics = append(clinics, resp)
	_, err = c.WriteFile(clinics)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *clinicRepo) Update(req *models.UpdateClinic) (*models.Clinic, error) {

	clinics, err := c.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, el := range clinics {
		if el.ClinicID == req.ClinicID {
			el.City = req.City
			el.Logo = req.Logo
			el.Name = req.Name
			_, err = c.WriteFile(clinics)
			if err != nil {
				return nil, err
			}
			return el, nil
		}

	}
	return &models.Clinic{}, nil
}

func (c *clinicRepo) GetByID(req *models.PrimeryKeyClinicID) (*models.Clinic, error) {

	clinics, err := c.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, el := range clinics {
		if el.ClinicID == req.ClinicID {
			return el, nil
		}

	}
	return &models.Clinic{}, nil
}

func (c *clinicRepo) Delete(req *models.PrimeryKeyClinicID) (string, error) {

	clinics, err := c.ReadAll()
	if err != nil {
		return "Bazada xatolik yuz berdi", err
	}

	for index, el := range clinics {
		if el.ClinicID == req.ClinicID {
			clinics = append(clinics[:index], clinics[index+1:]...)
			c.WriteFile(clinics)
			return "DELETED.....", nil
		}

	}
	return fmt.Sprintf("%s Shu IDga ega clinic topilmadi...", req.ClinicID), nil
}
