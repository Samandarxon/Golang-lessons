package model

type Category struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	ParentID  string `json:"parent_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CategoryCreate struct {
	Title     string `json:"title"`
	ParentID  string `json:"parent_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CategoryPrimaryKey struct {
	Id string `json:"id"`
}

type CategoryUpdate struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	ParentID  string `json:"parent_id"`
	UpdatedAt string `json:"updated_at"`
}

type GetListCategoryRequest struct {
	Offset int
	Limit  int
}

type GetListCategoryResponse struct {
	Count     int
	Categores []Category
}
