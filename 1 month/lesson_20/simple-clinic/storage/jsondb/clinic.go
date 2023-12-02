package jsondb

import (
	"errors"
	"os"

	"simple-clinic/config"
	"simple-clinic/models"
	"simple-clinic/pkg/helpers"
	"simple-clinic/storage"

	"github.com/google/uuid"
	"github.com/spf13/cast"
)

type clinicRepo struct {
	cfg  *config.Config
	file *os.File
}

func NewClinicRepo(cfg *config.Config, file *os.File) storage.ClinicRepoI {
	return &clinicRepo{cfg: cfg, file: file}
}

func (c *clinicRepo) Create(req *models.CreateClinic) (*models.Clinic, error) {

	data, err := helpers.Read(c.cfg.ClinicPath)
	if err != nil {
		return nil, err
	}

	var resp = &models.Clinic{
		ClinicID: uuid.New().String(),
		Name:     req.Name,
		Logo:     req.Logo,
		City:     req.City,
	}

	data = append(data, resp)

	err = helpers.Write(c.cfg.ClinicPath, data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *clinicRepo) GetByID(req *models.ClinicPrimaryKey) (*models.Clinic, error) {

	data, err := helpers.Read(c.cfg.ClinicPath)
	if err != nil {
		return nil, err
	}

	for _, clinicObj := range data {
		var clinic = cast.ToStringMap(clinicObj)
		if cast.ToString(clinic["clinic_id"]) == req.ID {

			return &models.Clinic{
				ClinicID: cast.ToString(clinic["clinic_id"]),
				Name:     cast.ToString(clinic["name"]),
				Logo:     cast.ToString(clinic["logo"]),
				City:     cast.ToString(clinic["city"]),
			}, nil
		}
	}

	return nil, errors.New("data is no rows")
}

func (c *clinicRepo) GetList(req *models.GetListClinicRequest) (*models.GetListClinicResponse, error) {
	data, err := helpers.Read(c.cfg.ClinicPath)
	if err != nil {
		return nil, err
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	// off: 100
	// lim: 10

	var resp = &models.GetListClinicResponse{}

	for i := req.Offset; i < req.Limit+req.Offset; i++ {
		if len(data) > i {
			var clinic = cast.ToStringMap(data[i])
			resp.Clinics = append(resp.Clinics, &models.Clinic{
				ClinicID: cast.ToString(clinic["clinic_id"]),
				Name:     cast.ToString(clinic["name"]),
				Logo:     cast.ToString(clinic["logo"]),
				City:     cast.ToString(clinic["city"]),
			})
		} else {
			break
		}
	}
	resp.Count = len(data)

	return resp, nil
}

func (c *clinicRepo) Update(req *models.UpdateClinic) (*models.Clinic, error) {

	data, err := helpers.Read(c.cfg.ClinicPath)
	if err != nil {
		return nil, err
	}

	for ind, clinicObj := range data {
		var clinic = cast.ToStringMap(clinicObj)

		if cast.ToString(clinic["clinic_id"]) == req.ClinicID {
			clinic["name"] = req.Name
			clinic["logo"] = req.Logo
			clinic["city"] = req.City

			data[ind] = clinic
			break
		}
	}

	err = helpers.Write(c.cfg.ClinicPath, data)
	if err != nil {
		return nil, err
	}

	return c.GetByID(&models.ClinicPrimaryKey{ID: req.ClinicID})
}

func (c *clinicRepo) Delete(req *models.ClinicPrimaryKey) error {

	data, err := helpers.Read(c.cfg.ClinicPath)
	if err != nil {
		return err
	}

	for ind, clinicObj := range data {
		var clinic = cast.ToStringMap(clinicObj)
		if cast.ToString(clinic["clinic_id"]) == req.ID {
			data = append(data[:ind], data[ind+1:]...)
			break
		}
	}

	err = helpers.Write(c.cfg.ClinicPath, data)
	if err != nil {
		return err
	}

	return nil
}
