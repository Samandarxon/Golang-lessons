package storage

import "github.com/asadbekGo/market_system/models"

type StorageI interface {
	Category() CategoryRepoI
	Product() ProductRepoI
	Remaining() RemainingRepoI
	Branch() BranchRepoI
	Coming() ComingRepoI
}

type CategoryRepoI interface {
	Create(req *models.CreateCategory) (*models.Category, error)
	GetByID(req *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Update(req *models.UpdateCategory) (int64, error)
	Delete(req *models.CategoryPrimaryKey) error
}
type ProductRepoI interface {
	Create(req *models.CreateProduct) (*models.Product, error)
	GetByID(req *models.ProductPrimaryKey) (*models.Product, error)
	GetList(req *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(req *models.UpdateProduct) (int64, error)
	Delete(req *models.ProductPrimaryKey) error
}
type RemainingRepoI interface {
	Create(req *models.CreateRemaining) (*models.Remaining, error)
	GetByID(req *models.RemainingPrimaryKey) (*models.Remaining, error)
	GetList(req *models.GetListRemainingRequest) (*models.GetListRemainingResponse, error)
	Update(req *models.UpdateRemaining) (int64, error)
	Delete(req *models.RemainingPrimaryKey) error
}
type BranchRepoI interface {
	Create(req *models.CreateBranch) (*models.Branch, error)
	GetByID(req *models.BranchPrimaryKey) (*models.Branch, error)
	GetList(req *models.GetListBranchRequest) (*models.GetListBranchResponse, error)
	Update(req *models.UpdateBranch) (int64, error)
	Delete(req *models.BranchPrimaryKey) error
}

type ComingRepoI interface {
	Create(req *models.CreateComing) (*models.Coming, error)
	GetByID(req *models.ComingPrimaryKey) (*models.Coming, error)
	GetList(req *models.GetListComingRequest) (*models.GetListComingResponse, error)
	Update(req *models.UpdateComing) (int64, error)
	Delete(req *models.ComingPrimaryKey) error
}
