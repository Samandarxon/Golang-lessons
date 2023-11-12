package postgres

import (
	"database/sql"
	"fmt"

	"example.com/m/config"
	"example.com/m/storage"
)

type StorePostgres struct {
	db      *sql.DB
	product *ProductRepo
}

func NewPostgressConnection(cfg *config.Config) (storage.StorageI, error) {
	connect := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	)

	db, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, err
	}
	// fmt.Println("db connect....", db)

	return &StorePostgres{
		db: db,
	}, err
}

func (s *StorePostgres) CloseDB() error {
	err := s.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (s *StorePostgres) Product() storage.ProductRepoI {
	if s.product == nil {
		s.product = NewproductRepoConnection(s.db)
	}

	return s.product
}
