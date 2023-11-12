package postgres

import (
	"database/sql"
	"fmt"

	"github.com/asadbekGo/market_system/models"
	"github.com/asadbekGo/market_system/pkg/helpers"
	"github.com/google/uuid"
)

type remainingRepo struct {
	db *sql.DB
}

func NewRemainingRepo(db *sql.DB) *remainingRepo {
	return &remainingRepo{
		db: db,
	}
}

func (r *remainingRepo) Create(req *models.CreateRemaining) (*models.Remaining, error) {

	var (
		remainingId = uuid.New().String()
		query       = `
			INSERT INTO "remaining"(
				"id",
				"branch_id",
				"product_id",
				"quantity",
				"price",
				"updated_at"
			) VALUES ($1, $2, $3,$4,$5, NOW())`
	)

	_, err := r.db.Exec(
		query,
		remainingId,
		helpers.NewNullString(req.Branch_id),
		helpers.NewNullString(req.Product_id),
		req.Quantity,
		req.Price,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(&models.RemainingPrimaryKey{Id: remainingId})
}

func (r *remainingRepo) GetByID(req *models.RemainingPrimaryKey) (*models.Remaining, error) {

	var (
		remaining models.Remaining
		query     = `
			SELECT
			"id",
			COALESCE(CAST("branch_id" AS VARCHAR), ''),
			COALESCE(CAST("product_id" AS VARCHAR), ''),
			"quantity",
			"price",
			"created_at",
			"updated_at"	
			FROM "remaining"
			WHERE "id" = $1
		`
	)

	err := r.db.QueryRow(query, req.Id).Scan(
		&remaining.Id,
		&remaining.Branch_id,
		&remaining.Product_id,
		&remaining.Quantity,
		&remaining.Price,
		&remaining.CreatedAt,
		&remaining.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &remaining, nil
}

func (r *remainingRepo) GetList(req *models.GetListRemainingRequest) (*models.GetListRemainingResponse, error) {
	var (
		resp   models.GetListRemainingResponse
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
		where += " AND price ILIKE" + " '%" + req.Search + "%'"
	}

	var query = `
		SELECT
			COUNT(*) OVER(),
			"id",
			"branch_id",
			"product_id",
			"quantity",
			"price", 
			"created_at",
			"updated_at"
		FROM "remaining"
	`

	query += where + sort + offset + limit
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			remaining  models.Remaining
			product_ID sql.NullString
			branch_ID  sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
			&remaining.Id,
			&branch_ID,
			&product_ID,
			&remaining.Quantity,
			&remaining.Price,
			&remaining.CreatedAt,
			&remaining.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		remaining.Branch_id = branch_ID.String
		remaining.Product_id = product_ID.String
		resp.Remainings = append(resp.Remainings, &remaining)
	}

	return &resp, nil
}

func (r *remainingRepo) Update(req *models.UpdateRemaining) (int64, error) {

	query := `
		UPDATE remaining
			SET
				"branch_id"=$2,
				"product_id"=$3,
				"quantity"=$4,
				"price"=$5, 
				"updated_at"=NOW()
		WHERE id = $1
	`
	result, err := r.db.Exec(
		query,
		req.Id,
		helpers.NewNullString(req.Branch_id),
		helpers.NewNullString(req.Product_id),
		req.Quantity,
		req.Price,
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

func (r *remainingRepo) Delete(req *models.RemainingPrimaryKey) error {
	_, err := r.db.Exec("DELETE FROM remaining WHERE id = $1", req.Id)
	return err
}
