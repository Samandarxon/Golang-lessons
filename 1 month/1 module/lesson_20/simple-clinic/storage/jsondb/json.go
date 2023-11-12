package jsondb

import (
	"os"
	"simple-clinic/config"
	"simple-clinic/storage"
)

type Store struct {
	clinic  storage.ClinicRepoI
	branch  storage.BranchRepoI
	cashier storage.CashierRepoI
}

func NewJsonDbConnection(cfg *config.Config) (storage.StorageI, error) {

	clinicFile, err := os.Open(cfg.ClinicPath)
	if err != nil {
		return nil, err
	}

	branchFile, err := os.Open(cfg.BranchPath)
	if err != nil {
		return nil, err
	}
	cashierFile, err := os.Open(cfg.CashierPath)
	if err != nil {
		return nil, err
	}

	return &Store{
		clinic:  NewClinicRepo(cfg, clinicFile),
		branch:  NewBranchRepo(cfg, branchFile),
		cashier: NewCashierRepo(cfg, cashierFile),
	}, nil
}

func (s Store) Clinic() storage.ClinicRepoI {
	return s.clinic
}

func (s Store) Branch() storage.BranchRepoI {
	return s.branch
}
func (s Store) Cashier() storage.CashierRepoI {
	return s.cashier
}
