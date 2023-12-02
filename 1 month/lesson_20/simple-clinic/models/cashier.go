package models

type CashierPrimaryKey struct {
	ID string `json:"id"`
}

type CreateCashier struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone"`
	BranchID string `json:"branch_id"`
}

type Cashier struct {
	CashierID string `json:"cashier_id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Phone     string `json:"phone"`
	BranchID  string `json:"branch_id"`
}

type UpdateCashier struct {
	CashierID string `json:"cashier_id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Phone     string `json:"phone"`
	BranchID  string `json:"branch_id"`
}

type GetListCashierRequest struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Search string `json:"search"`
}

type GetListCashierResponse struct {
	Count    int        `json:"count"`
	Cashiers []*Cashier `json:"Cashiers"`
}
