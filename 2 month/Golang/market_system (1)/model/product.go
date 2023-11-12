package model

type Product struct {
	Id          string
	Name        string
	Barcode     string
	Price       float64
	Category_id string
	Image_url   string
	Updated_at  string
}

type ProductCreate struct {
	Name        string
	Barcode     string
	Price       float64
	Category_id string
	Image_url   string
}

type ProductPrimaryKey struct {
	Id string
}

type ProductUpdate struct {
	Id          string
	Name        string
	Barcode     string
	Price       float64
	Category_id string
	Image_url   string
}

type GetListProductRequest struct {
	Offset int
	Limit  int
}

type GetListProductResponse struct {
	Count    int
	Products []Product
}
