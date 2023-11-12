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

type cashierRepo struct {
	cfg  *config.Config
	file *os.File
}

func NewCashierRepo(cfg *config.Config, file *os.File) storage.CashierRepoI {
	return &cashierRepo{cfg: cfg, file: file}
}

func (c *cashierRepo) Create(req *models.CreateCashier) (*models.Cashier, error) {

	data, err := helpers.Read(c.cfg.CashierPath)
	if err != nil {
		return nil, err
	}

	var resp = &models.Cashier{
		CashierID: uuid.New().String(),
		Name:      req.Name,
		Surname:   req.Surname,
		Phone:     req.Phone,
		BranchID:  req.BranchID,
	}

	data = append(data, resp)

	err = helpers.Write(c.cfg.CashierPath, data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *cashierRepo) GetByID(req *models.CashierPrimaryKey) (*models.Cashier, error) {

	data, err := helpers.Read(c.cfg.CashierPath)
	if err != nil {
		return nil, err
	}

	for _, cashierObj := range data {
		var cashier = cast.ToStringMap(cashierObj)
		if cast.ToString(cashier["cashier_id"]) == req.ID {

			return &models.Cashier{
				CashierID: cast.ToString(cashier["cashier_id"]),
				BranchID:  cast.ToString(cashier["branch_id"]),
				Name:      cast.ToString(cashier["name"]),
				Surname:   cast.ToString(cashier["surname"]),
				Phone:     cast.ToString(cashier["phone"]),
			}, nil
		}
	}

	return nil, errors.New("data is no rows")
}

func (c *cashierRepo) GetList(req *models.GetListCashierRequest) (*models.GetListCashierResponse, error) {
	data, err := helpers.Read(c.cfg.CashierPath)
	if err != nil {
		return nil, err
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	// off: 100
	// lim: 10

	var resp = &models.GetListCashierResponse{}

	for i := req.Offset; i < req.Limit+req.Offset; i++ {
		if len(data) > i {
			var cashier = cast.ToStringMap(data[i])
			resp.Cashiers = append(resp.Cashiers, &models.Cashier{
				CashierID: cast.ToString(cashier["cashier_id"]),
				BranchID:  cast.ToString(cashier["branch_id"]),
				Name:      cast.ToString(cashier["name"]),
				Surname:   cast.ToString(cashier["surname"]),
				Phone:     cast.ToString(cashier["phone"]),
			})
		} else {
			break
		}
	}
	resp.Count = len(data)

	return resp, nil
}

func (c *cashierRepo) Update(req *models.UpdateCashier) (*models.Cashier, error) {

	data, err := helpers.Read(c.cfg.CashierPath)
	if err != nil {
		return nil, err
	}

	for ind, cashierObj := range data {
		var cashier = cast.ToStringMap(cashierObj)

		if cast.ToString(cashier["cashier_id"]) == req.CashierID {
			cashier["cashier_id"] = req.CashierID
			cashier["branch_id"] = req.BranchID
			cashier["name"] = req.Name
			cashier["surname"] = req.Surname
			cashier["phone"] = req.Phone

			data[ind] = cashier
			break
		}
	}

	err = helpers.Write(c.cfg.CashierPath, data)
	if err != nil {
		return nil, err
	}

	return c.GetByID(&models.CashierPrimaryKey{ID: req.CashierID})
}

func (c *cashierRepo) Delete(req *models.CashierPrimaryKey) error {

	data, err := helpers.Read(c.cfg.CashierPath)
	if err != nil {
		return err
	}

	for ind, cashierObj := range data {
		var cashier = cast.ToStringMap(cashierObj)
		if cast.ToString(cashier["cashier_id"]) == req.ID {
			data = append(data[:ind], data[ind+1:]...)
			break
		}
	}

	err = helpers.Write(c.cfg.CashierPath, data)
	if err != nil {
		return err
	}

	return nil
}
