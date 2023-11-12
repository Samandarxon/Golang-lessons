package models

type RemainingPrimaryKey struct {
	Id string `json:"id"`
}

type CreateRemaining struct {
	Branch_id  string  `json:"branch_id"`
	Product_id string  `json:"product_id"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
}

type Remaining struct {
	Id         string  `json:"id"`
	Branch_id  string  `json:"branch_id"`
	Product_id string  `json:"product_id"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type UpdateRemaining struct {
	Id         string  `json:"id"`
	Branch_id  string  `json:"branch_id"`
	Product_id string  `json:"product_id"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
}

type GetListRemainingRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}

type GetListRemainingResponse struct {
	Count      string       `json:"count"`
	Remainings []*Remaining `json:"products"`
}
