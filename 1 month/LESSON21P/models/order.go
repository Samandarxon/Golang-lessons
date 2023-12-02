package models

type Order struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Sum       int    `json:"sum"`
	SumCount  int    `json:"sum_count"`
	DateTime  string `json:"date_time"`
	Status    string `json:"status"`
	ProductId string `json:"product_id"`
}
type CreateOrder struct {
	UserId    string `json:"user_id"`
	Sum       int    `json:"sum"`
	SumCount  int    `json:"sum_count"`
	DateTime  string `json:"date_time"`
	Status    string `json:"status"`
	ProductId string `json:"product_id"`
}

type PrimeryKeyOrder struct {
	Id string `json:"id"`
}

type UpadetOrder struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Sum       int    `json:"sum"`
	SumCount  int    `json:"sum_count"`
	DateTime  string `json:"date_time"`
	Status    string `json:"status"`
	ProductId string `json:"product_id"`
}

type GetListOrderRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetListOrderResponse struct {
	Count  int      `json:"count"`
	Orders []*Order `json:"users"`
}
