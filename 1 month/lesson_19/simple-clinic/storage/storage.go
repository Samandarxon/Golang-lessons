package storage

import "simple-clinic/models"

type StorageI interface {
	Clinic() ClinicRepoI
	BranchOffice() BranchOfficeRepoI
}

type ClinicRepoI interface {
	Create(req *models.CreateClinic) (*models.Clinic, error)
	GetByID(req *models.PrimeryKeyClinicID) (*models.Clinic, error)
	Update(req *models.UpdateClinic) (*models.Clinic, error)
	Delete(req *models.PrimeryKeyClinicID) (string, error)
}

type BranchOfficeRepoI interface {
	Create(req *models.CreateBranchOffice) (*models.BranchOffice, error)
	GetByID(req *models.PrimeryKeyBranchOfficeID) (*models.BranchOffice, error)
	Update(req *models.UpdateBranchOffice) (*models.BranchOffice, error)
	Delete(req *models.PrimeryKeyBranchOfficeID) (string, error)
}
