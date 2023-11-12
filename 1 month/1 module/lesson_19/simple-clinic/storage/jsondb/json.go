package jsondb

import (
	"os"
	"simple-clinic/config"
	"simple-clinic/storage"
)

type Store struct {
	clinic       storage.ClinicRepoI
	branchOffice storage.BranchOfficeRepoI
}

func NewJsonDbConnection(cfg *config.Config) (storage.StorageI, error) {

	clinicFile, err := os.Open(cfg.ClinicPath)
	if err != nil {
		return nil, err
	}

	return &Store{
		clinic:       NewClinicRepo(cfg, clinicFile),
		branchOffice: NewBranchOfficeRepo(cfg, clinicFile),
	}, nil
}

func (s Store) Clinic() storage.ClinicRepoI {
	return s.clinic
}
func (s Store) BranchOffice() storage.BranchOfficeRepoI {
	return s.branchOffice
}
