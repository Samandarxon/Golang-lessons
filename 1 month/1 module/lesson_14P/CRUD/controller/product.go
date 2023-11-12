package controller

import (
	"fmt"
	"time"

	"app/models"
)

type ProductController struct {
	Products []models.Product
}

func NewProductController(products []models.Product) ProductController {
	return ProductController{
		Products: products,
	}
}

func (p *ProductController) CreateProduct(req models.ProductCreate) models.Product {
	var product = models.Product{
		Id:        req.Id,
		Name:      req.Name,
		Price:     req.Price,
		Quantity:  req.Quantity,
		CreatedAt: time.Now().Format("2006-01-01"),
		UpdatedAt: time.Now().Format("2006-01-01"),
	}

	p.Products = append(p.Products, product)

	return product
}

func (p *ProductController) ProductGetList(req models.ProductGetListRequest) models.ProductGetListResponse {
	var products models.ProductGetListResponse

	products.Count = len(p.Products)
	products.Products = p.Products

	return products
}

func (p *ProductController) ProdcutGetById(req models.ProductPrimaryKey) models.Product {
	for _, v := range p.Products {
		if v.Id == req.Id {
			return v
		}
	}

	return models.Product{}
}

func (p *ProductController) ProductUpdate(req models.ProductUpdate) models.Product {
	for index, el := range p.Products {
		if el.Id == req.Id {
			p.Products[index].Price = req.Price
			p.Products[index].Quantity = req.Quantity
			return p.Products[index]
		}
	}
	// fmt.Println(p.Products)
	return models.Product{}
}

func (p *ProductController) ProductDelete(req models.ProductPrimaryKey) string {
	for index, el := range p.Products {
		if el.Id == req.Id {
			// fmt.Println("del", p.Products[index+1:])
			p.Products = append(p.Products[:index], p.Products[index+1:]...)
			return "Deleted"
		}
	}
	return fmt.Sprintf("Bu %s ID-ga ega product yo'q", req.Id.String())
}
