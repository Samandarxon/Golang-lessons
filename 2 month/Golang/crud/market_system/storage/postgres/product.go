package postgres

import (
	"database/sql"
	"fmt"

	"github.com/asadbekGo/market_system/models"
	"github.com/asadbekGo/market_system/pkg/helpers"
	"github.com/google/uuid"
)

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(req *models.CreateProduct) (*models.Product, error) {

	var (
		productId = uuid.New().String()
		query     = `
			INSERT INTO "product"(
				"id",
				"name", 
				"price", 
				"category_id", 
				"updated_at"
			) VALUES ($1, $2, $3,$4, NOW())`
	)

	_, err := r.db.Exec(
		query,
		productId,
		req.Name,
		req.Price,
		helpers.NewNullString(req.Category_id),
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(&models.ProductPrimaryKey{Id: productId})
}

func (r *productRepo) GetByID(req *models.ProductPrimaryKey) (*models.Product, error) {

	var (
		product models.Product
		query   = `
			SELECT
			"id",
			"name", 
			"price", 
			COALESCE(CAST("category_id" AS VARCHAR), ''),
			"created_at",
			"updated_at"	
			FROM "product"
			WHERE "id" = $1
		`
	)

	err := r.db.QueryRow(query, req.Id).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Category_id,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepo) GetList(req *models.GetListProductRequest) (*models.GetListProductResponse, error) {
	var (
		resp   models.GetListProductResponse
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		sort   = " ORDER BY created_at DESC"
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if len(req.Search) > 0 {
		where += " AND name ILIKE" + " '%" + req.Search + "%'"
	}

	var query = `
		SELECT
			COUNT(*) OVER(),
			"id",
			"name", 
			"price", 
			"category_id", 
			"created_at",
			"updated_at"
		FROM "product"
	`

	query += where + sort + offset + limit
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			product    models.Product
			categoryID sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
			&product.Id,
			&product.Name,
			&product.Price,
			&categoryID,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		product.Category_id = categoryID.String
		resp.Products = append(resp.Products, &product)
	}

	return &resp, nil
}

func (r *productRepo) Update(req *models.UpdateProduct) (int64, error) {

	query := `
		UPDATE product
			SET
				name=$2, 
				price=$3, 
				category_id=$4, 
				updated_at=NOW()
		WHERE id = $1
	`
	result, err := r.db.Exec(
		query,
		req.Id,
		req.Name,
		req.Price,
		helpers.NewNullString(req.Category_id),
	)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *productRepo) Delete(req *models.ProductPrimaryKey) error {
	_, err := r.db.Exec("DELETE FROM product WHERE id = $1", req.Id)
	return err
}
