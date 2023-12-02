package models

import "github.com/google/uuid"

type Category struct {
	Id        uuid.UUID
	Name      string
	CreatedAt string
	UpdatedAt string
}

type CreateCategoryRequest struct {
	Name string
}
type CreateCategoryRespons struct {
	Category
}

type CategoryGetListRequest struct {
	OffSet int
	Limit  int
}

type CategoryGetListRespons struct {
	Count     int
	Categorys []Category
}

type CategoryPrimaryKey struct {
	Id uuid.UUID
}
type CategoryUpdateRequest struct {
	Id   uuid.UUID
	Name string
}

type CategoryUpdateResponse struct {
	Category
}
