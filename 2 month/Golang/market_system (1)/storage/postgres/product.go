package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"example.com/m/model"
	"github.com/google/uuid"
)

type ProductRepo struct {
	db *sql.DB
}

func NewproductRepoConnection(db *sql.DB) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (p *ProductRepo) Create(req model.ProductCreate) (*model.Product, error) {
	var (
		productId = uuid.New().String()
		Product   = model.Product{
			// Id:          productId,
			Name:        req.Name,
			Barcode:     req.Barcode,
			Price:       req.Price,
			Category_id: req.Category_id,
			Image_url:   req.Image_url,
		}

		query = `
			INSERT INTO products(
				name, 
				barcode, 
				price,
				category_id,
				image_url,
				updated_at
			) 
			VALUES ($1, $2, $3, $4, $5, NOW())`
	)

	_, err := p.db.Exec(
		query,
		Product.Name,
		Product.Barcode,
		Product.Price,
		productId,
		Product.Image_url,
	)
	if err != nil {
		return nil, err
	}

	return &Product, nil
}
func (p *ProductRepo) Update(req model.ProductUpdate) (*model.Product, error) {
	var (
		Product = model.Product{
			Id:          req.Id,
			Name:        req.Name,
			Barcode:     req.Barcode,
			Price:       req.Price,
			Category_id: req.Category_id,
			Image_url:   req.Image_url,
		}

		query = `
			UPDATE products SET
				name=$1, 
				barcode=$2, 
				price=$3,
				image_url=$4,
				updated_at=NOW()
			WHERE id = $5
			`
	)
	result, err := p.db.Exec(
		query,
		Product.Name,
		Product.Barcode,
		Product.Price,
		Product.Image_url,
		Product.Id,
	)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)

	return &Product, nil
}
func (p *ProductRepo) GetList() ([]model.Product, error) {
	query := `SELECT 
			*
	FROM products`

	rows, err := p.db.Query(
		query,
	)

	if err != nil {
		return nil, err
	}
	var products []model.Product
	for rows.Next() {
		var (
			product    model.Product
			categoryID sql.NullString
		)
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Barcode,
			&product.Price,
			&categoryID,
			&product.Image_url,
			&product.Updated_at,
		)
		if err != nil {
			log.Fatal(err)
		}
		// log.Println(id, title)
		product.Category_id = categoryID.String
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(products)

	return products, nil
}
func (p *ProductRepo) GetId(req model.ProductPrimaryKey) (*model.Product, error) {

	var (
		product model.Product
		query   = `
			SELECT
				id,
				name,
				barcode,
				price,
				COALESCE(CAST("category_id" AS VARCHAR), ''),
				image_url,
				updated_at
			FROM "products"
			WHERE "id" = $1
		`
	)

	// query row function bitta row olish uchun ishlatiladi
	err := p.db.QueryRow(query, req.Id).Scan(
		&product.Id,
		&product.Name,
		&product.Barcode,
		&product.Price,
		&product.Category_id,
		&product.Image_url,
		&product.Updated_at,
	)

	if err != nil {
		log.Println("db query row function:", err.Error())
		return nil, err
	}

	fmt.Printf("Product: %+v\n", product)

	// fmt.Println(products)

	return &product, nil
}

func (p *ProductRepo) Delete(req model.ProductPrimaryKey) (sql.Result, error) {
	query := `DELETE FROM products WHERE id = $1`

	result, err := p.db.Exec(
		query,
		req.Id,
	)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)

	return result, nil
}
