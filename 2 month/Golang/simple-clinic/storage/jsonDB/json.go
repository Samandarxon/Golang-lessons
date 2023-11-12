package jsonDB

import (
	"app/config"
	"app/storage"
)

type Store struct {
	user     storage.UserRepoI
	category storage.CategoryRepoI
	order    storage.OrederRepoI
	product  storage.ProductRepoI
}

func NewJsonDbConnection(cfg *config.Config_DB) (storage.StorageI, error) {

	return &Store{
		user:     NewUserRepo(cfg),
		category: NewCategoryRepo(cfg),
		order:    NewOrderRepo(cfg),
		product:  NewProductRepo(cfg),
	}, nil
}

func (s Store) User() storage.UserRepoI {
	return s.user
}
func (s Store) Category() storage.CategoryRepoI {
	return s.category
}
func (s Store) Order() storage.OrederRepoI {
	return s.order
}
func (s Store) Product() storage.ProductRepoI {
	return s.product
}
