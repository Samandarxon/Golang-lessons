package storage

import (
	"database/sql"

	"example.com/m/model"
)

type StorageI interface {
	CloseDB() error
	Product() ProductRepoI
}

type ProductRepoI interface {
	Create(model.ProductCreate) (*model.Product, error)
	Update(model.ProductUpdate) (*model.Product, error)
	Delete(model.ProductPrimaryKey) (sql.Result, error)
	GetList() ([]model.Product, error)
	GetId(req model.ProductPrimaryKey) (*model.Product, error)
}
type CategoryRepoI interface {
	Create(model.CategoryCreate) (*model.Category, error)
	Update(model.CategoryUpdate) (*model.Category, error)
	Delete(model.CategoryPrimaryKey) (sql.Result, error)
	GetList() ([]model.Category, error)
	GetId(req model.CategoryPrimaryKey) (*model.Category, error)
}
