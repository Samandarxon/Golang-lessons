package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"example.com/m/model"
	"github.com/google/uuid"
)

type CategoryRepo struct {
	db *sql.DB
}

func NewcategoryRepoConnection(db *sql.DB) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (p *CategoryRepo) Create(req model.CategoryCreate) (*model.Category, error) {
	var (
		categoryId = uuid.New().String()
		Category   = model.Category{
			// Id:          categoryId,     string `json:"title"`
			Title:     req.Title,
			ParentID:  req.ParentID,
			CreatedAt: req.CreatedAt,
			UpdatedAt: req.UpdatedAt,
		}

		query = `
			INSERT INTO category(
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
		Category.Title,
		Category.ParentID,
		Category.CreatedAt,
		Category.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &Category, nil
}
func (p *CategoryRepo) Update(req model.CategoryUpdate) (*model.Category, error) {
	var (
		Category = model.Category{
			Id:          req.Id,
			Name:        req.Name,
			Barcode:     req.Barcode,
			Price:       req.Price,
			Category_id: req.Category_id,
			Image_url:   req.Image_url,
		}

		query = `
			UPDATE category SET
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
		Category.Name,
		Category.Barcode,
		Category.Price,
		Category.Image_url,
		Category.Id,
	)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)

	return &Category, nil
}
func (p *CategoryRepo) GetList() ([]model.Category, error) {
	query := `SELECT 
			*
	FROM category`

	rows, err := p.db.Query(
		query,
	)

	if err != nil {
		return nil, err
	}
	var categorys []model.Category
	for rows.Next() {
		var (
			category   model.Category
			categoryID sql.NullString
		)
		err := rows.Scan(
			&category.Id,
			&category.Name,
			&category.Barcode,
			&category.Price,
			&categoryID,
			&category.Image_url,
			&category.Updated_at,
		)
		if err != nil {
			log.Fatal(err)
		}
		// log.Println(id, title)
		category.Category_id = categoryID.String
		categorys = append(categorys, category)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(categorys)

	return categorys, nil
}
func (p *CategoryRepo) GetId(req model.CategoryPrimaryKey) (*model.Category, error) {

	var (
		category model.Category
		query    = `
			SELECT
				id,
				name,
				barcode,
				price,
				COALESCE(CAST("category_id" AS VARCHAR), ''),
				image_url,
				updated_at
			FROM "category"
			WHERE "id" = $1
		`
	)

	// query row function bitta row olish uchun ishlatiladi
	err := p.db.QueryRow(query, req.Id).Scan(
		&category.Id,
		&category.Name,
		&category.Barcode,
		&category.Price,
		&category.Category_id,
		&category.Image_url,
		&category.Updated_at,
	)

	if err != nil {
		log.Println("db query row function:", err.Error())
		return nil, err
	}

	fmt.Printf("Category: %+v\n", category)

	// fmt.Println(categorys)

	return &category, nil
}

func (p *CategoryRepo) Delete(req model.CategoryPrimaryKey) (sql.Result, error) {
	query := `DELETE FROM category WHERE id = $1`

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
