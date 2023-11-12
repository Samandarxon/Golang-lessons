package models

type BranchPrimaryKey struct {
	Id string `json:"id"`
}

type CreateBranch struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Branch struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateBranch struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type GetListBranchRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}

type GetListBranchResponse struct {
	Count   string    `json:"count"`
	Branchs []*Branch `json:"branchs"`
}
