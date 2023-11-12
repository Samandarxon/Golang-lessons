package models

type Category struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CreateCategory struct {
	Name string `json:"name"`
}
type CreateCategoryRespons struct {
	Category
}

type CategoryGetListRequest struct {
	OffSet int
	Limit  int
}

type CategoryGetListRespons struct {
	Count     int        `json:"count"`
	Categorys []Category `json:"categorys"`
}

type CategoryPrimaryKey struct {
	Id string `json:"id"`
}
type CategoryUpdateRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
