package storage

import "app/models"

type StorageI interface {
	User() UserRepoI
	Category() CategoryRepoI
	Order() OrederRepoI
	Product() ProductRepoI
}

type UserRepoI interface {
	Create(req *models.CreateUser) (*models.User, error)
	Update(req *models.UpadetUser) (*models.User, error)
	GetByID(req *models.PrimeryKeyUser) (*models.User, error)
	GetList(req *models.GetListUserRequest) (*models.GetListUserResponse, error)
	Delete(req *models.PrimeryKeyUser) error
}

type CategoryRepoI interface {
	Create(req *models.CreateCategory) (*models.Category, error)
	// Update(req *models.UpadetCategory) (*models.Category, error)
	// GetByID(req *models.PrimeryKeyCategory) (*models.Category, error)
	// GetList(req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	// Delete(req *models.PrimeryKeyCategory) error
}

type OrederRepoI interface {
	Create(req *models.CreateOrder) (*models.Order, error)
	Update(req *models.UpadetOrder) (*models.Order, error)
	GetByID(req *models.PrimeryKeyOrder) (*models.Order, error)
	GetList(req *models.GetListOrderRequest) (*models.GetListOrderResponse, error)
	Delete(req *models.PrimeryKeyOrder) error
}

type ProductRepoI interface {
	Create(req *models.CreateProduct) (*models.Product, error)
	Update(req *models.UpadetProduct) (*models.Product, error)
	GetByID(req *models.PrimeryKeyProduct) (*models.Product, error)
	GetList(req *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Delete(req *models.PrimeryKeyProduct) error
}
