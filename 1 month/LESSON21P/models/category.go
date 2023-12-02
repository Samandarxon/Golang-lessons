package models

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type CreateCategory struct {
	Name string `json:"name"`
}

type PrimeryKeyCategory struct {
	Id string `json:"id"`
}

type UpadetCategory struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetListCategoryRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetListCategoryResponse struct {
	Count     int         `json:"count"`
	Categorys []*Category `json:"users"`
}
