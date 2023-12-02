package models

type BranchPrimaryKey struct {
	ID string `json:"id"`
}

type CreateBranch struct {
	ClinicID string `json:"clinic_id"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type Branch struct {
	BranchID string `json:"branch_id"`
	ClinicID string `json:"clinic_id"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type UpdateBranch struct {
	BranchID string `json:"branch_id"`
	ClinicID string `json:"clinic_id"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type GetListBranchRequest struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Search string `json:"search"`
}

type GetListBranchResponse struct {
	Count   int       `json:"count"`
	Branchs []*Branch `json:"branchs"`
}
