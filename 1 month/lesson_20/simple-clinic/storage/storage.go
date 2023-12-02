package storage

import "simple-clinic/models"

type StorageI interface {
	Clinic() ClinicRepoI
	Branch() BranchRepoI
	Cashier() CashierRepoI
}

type ClinicRepoI interface {
	Create(req *models.CreateClinic) (*models.Clinic, error)
	GetByID(req *models.ClinicPrimaryKey) (*models.Clinic, error)
	GetList(req *models.GetListClinicRequest) (*models.GetListClinicResponse, error)
	Update(req *models.UpdateClinic) (*models.Clinic, error)
	Delete(req *models.ClinicPrimaryKey) error
}

type BranchRepoI interface {
	Create(req *models.CreateBranch) (*models.Branch, error)
	GetByID(req *models.BranchPrimaryKey) (*models.Branch, error)
	GetList(req *models.GetListBranchRequest) (*models.GetListBranchResponse, error)
	Update(req *models.UpdateBranch) (*models.Branch, error)
	Delete(req *models.BranchPrimaryKey) error
}
type CashierRepoI interface {
	Create(req *models.CreateCashier) (*models.Cashier, error)
	GetByID(req *models.CashierPrimaryKey) (*models.Cashier, error)
	GetList(req *models.GetListCashierRequest) (*models.GetListCashierResponse, error)
	Update(req *models.UpdateCashier) (*models.Cashier, error)
	Delete(req *models.CashierPrimaryKey) error
}
