package jsonDB

import (
	"app/config"
	"app/models"
	"app/pkg/helpers"
	"app/storage"
	"errors"
	"fmt"

	"github.com/spf13/cast"
)

type productRepo struct {
	productFile string
}

func NewProductRepo(cfg *config.Config) storage.ProductRepoI {

	return &productRepo{productFile: cfg.FilePath + cfg.File.ProductFile}
}

func (u *productRepo) Create(req *models.CreateProduct) (*models.Product, error) {

	data, err := helpers.Read(u.productFile)
	if err != nil {
		return nil, err
	}

	var resp = &models.Product{
		Name:         req.Name,
		Price:        req.Price,
		Discount:     req.Discount,
		DiscountType: req.DiscountType,
		CategoryID:   req.CategoryID,
	}

	data = append(data, resp)

	err = helpers.Write(u.productFile, data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *productRepo) GetByID(req *models.PrimeryKeyProduct) (*models.Product, error) {

	data, err := helpers.Read(c.productFile)
	if err != nil {
		return nil, err
	}

	for _, productObj := range data {
		var product = cast.ToStringMap(productObj)
		if cast.ToString(product["id"]) == req.Id {

			return &models.Product{
				Id:           cast.ToString(product["id"]),
				Name:         cast.ToString(product["name"]),
				Price:        cast.ToInt(product["price"]),
				Discount:     cast.ToInt(product["discount"]),
				DiscountType: cast.ToString(product["discount_type"]),
				CategoryID:   cast.ToString(product["category_id"]),
			}, nil
		}
	}

	return nil, errors.New("data is no rows")
}

func (c *productRepo) GetList(req *models.GetListProductRequest) (*models.GetListProductResponse, error) {
	data, err := helpers.Read(c.productFile)
	if err != nil {
		return nil, err
	}
	// fmt.Println(data)
	if req.Limit == 0 {
		req.Limit = 10
	}

	var resp = &models.GetListProductResponse{}

	for i := req.Offset; i < req.Limit+req.Offset; i++ {
		if len(data) > i {
			var product = cast.ToStringMap(data[i])
			fmt.Println(i+1, product)
			resp.Products = append(resp.Products, &models.Product{
				Id:           cast.ToString(product["id"]),
				Name:         cast.ToString(product["name"]),
				Price:        cast.ToInt(product["price"]),
				Discount:     cast.ToInt(product["discount"]),
				DiscountType: cast.ToString(product["discount_type"]),
				CategoryID:   cast.ToString(product["category_id"]),
			})
		} else {
			break
		}
	}
	resp.Count = len(data)

	return resp, nil
}

func (c *productRepo) Update(req *models.UpadetProduct) (*models.Product, error) {

	data, err := helpers.Read(c.productFile)
	if err != nil {
		return nil, err
	}

	for _, productObj := range data {
		var product = cast.ToStringMap(productObj)

		if cast.ToString(product["id"]) == req.Id {
			product["id"] = req.Id
			product["name"] = req.Name
			product["price"] = req.Price
			product["discount"] = req.Discount
			product["discount_type"] = req.DiscountType
			product["category_id"] = req.CategoryID

			// data[ind] = product
			break
		}
	}

	err = helpers.Write(c.productFile, data)
	if err != nil {
		return nil, err
	}

	return c.GetByID(&models.PrimeryKeyProduct{Id: req.Id})
}

func (c *productRepo) Delete(req *models.PrimeryKeyProduct) error {

	data, err := helpers.Read(c.productFile)
	if err != nil {
		return err
	}

	for ind, productObj := range data {
		var product = cast.ToStringMap(productObj)
		if cast.ToString(product["id"]) == req.Id {
			data = append(data[:ind], data[ind+1:]...)
			break
		}
	}

	err = helpers.Write(c.productFile, data)
	if err != nil {
		return err
	}

	return nil
}
