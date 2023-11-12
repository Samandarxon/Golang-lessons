package models

type Product struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Discount     int    `json:"discount"`
	DiscountType string `json:"discount_type"`
	CategoryID   string `json:"category_id"`
}
type CreateProduct struct {
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Discount     int    `json:"discount"`
	DiscountType string `json:"discount_type"`
	CategoryID   string `json:"category_id"`
}

type PrimeryKeyProduct struct {
	Id string `json:"id"`
}

type UpadetProduct struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Discount     int    `json:"discount"`
	DiscountType string `json:"discount_type"`
	CategoryID   string `json:"category_id"`
}

type GetListProductRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetListProductResponse struct {
	Count    int        `json:"count"`
	Products []*Product `json:"users"`
}
