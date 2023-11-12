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

type branchRepo struct {
	cfg  *config.Config
	file *os.File
}

func NewBranchRepo(cfg *config.Config, file *os.File) storage.BranchRepoI {
	return &branchRepo{cfg: cfg, file: file}
}

func (c *branchRepo) Create(req *models.CreateBranch) (*models.Branch, error) {

	data, err := helpers.Read(c.cfg.BranchPath)
	if err != nil {
		return nil, err
	}

	var resp = &models.Branch{
		BranchID: uuid.New().String(),
		ClinicID: req.ClinicID,
		Address:  req.Address,
		Phone:    req.Phone,
	}

	data = append(data, resp)

	err = helpers.Write(c.cfg.BranchPath, data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *branchRepo) GetByID(req *models.BranchPrimaryKey) (*models.Branch, error) {

	data, err := helpers.Read(c.cfg.BranchPath)
	if err != nil {
		return nil, err
	}

	for _, branchObj := range data {
		var branch = cast.ToStringMap(branchObj)
		if cast.ToString(branch["branch_id"]) == req.ID {

			return &models.Branch{
				BranchID: cast.ToString(branch["branch_id"]),
				ClinicID: cast.ToString(branch["clinic_id"]),
				Address:  cast.ToString(branch["address"]),
				Phone:    cast.ToString(branch["phone"]),
			}, nil
		}
	}

	return nil, errors.New("data is no rows")
}

func (c *branchRepo) GetList(req *models.GetListBranchRequest) (*models.GetListBranchResponse, error) {
	data, err := helpers.Read(c.cfg.BranchPath)
	if err != nil {
		return nil, err
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	// off: 100
	// lim: 10

	var resp = &models.GetListBranchResponse{}

	for i := req.Offset; i < req.Limit+req.Offset; i++ {
		if len(data) > i {
			var branch = cast.ToStringMap(data[i])
			resp.Branchs = append(resp.Branchs, &models.Branch{
				BranchID: cast.ToString(branch["branch_id"]),
				ClinicID: cast.ToString(branch["clinic_id"]),
				Address:  cast.ToString(branch["address"]),
				Phone:    cast.ToString(branch["phone"]),
			})
		} else {
			break
		}
	}
	resp.Count = len(data)

	return resp, nil
}

func (c *branchRepo) Update(req *models.UpdateBranch) (*models.Branch, error) {

	data, err := helpers.Read(c.cfg.BranchPath)
	if err != nil {
		return nil, err
	}

	for ind, branchObj := range data {
		var branch = cast.ToStringMap(branchObj)

		if cast.ToString(branch["branch_id"]) == req.BranchID {
			branch["clinic_id"] = req.ClinicID
			branch["address"] = req.Address
			branch["phone"] = req.Phone

			data[ind] = branch
			break
		}
	}

	err = helpers.Write(c.cfg.BranchPath, data)
	if err != nil {
		return nil, err
	}

	return c.GetByID(&models.BranchPrimaryKey{ID: req.BranchID})
}

func (c *branchRepo) Delete(req *models.BranchPrimaryKey) error {

	data, err := helpers.Read(c.cfg.BranchPath)
	if err != nil {
		return err
	}

	for ind, branchObj := range data {
		var branch = cast.ToStringMap(branchObj)
		if cast.ToString(branch["branch_id"]) == req.ID {
			data = append(data[:ind], data[ind+1:]...)
			break
		}
	}

	err = helpers.Write(c.cfg.BranchPath, data)
	if err != nil {
		return err
	}

	return nil
}
