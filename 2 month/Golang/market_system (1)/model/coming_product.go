package model

type ComingProduct struct {
	Id         string
	Category   string
	Product    string
	Count      int
	Price      float64
	TotalPirce float64
}

type ComingProductPrimaryKey struct {
	Id string
}

type ComingProductUpdate struct {
	Id         string
	Category   string
	Product    string
	Count      int
	Price      float64
	TotalPirce float64
}

type GetListComingProductRequest struct {
	Offset int
	Limit  int
}

type GetListComingProductResponse struct {
	Count          int
	ComingProducts []ComingProduct
}
