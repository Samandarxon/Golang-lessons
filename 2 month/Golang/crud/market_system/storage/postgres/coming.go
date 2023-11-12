package postgres

import (
	"database/sql"
	"fmt"

	"github.com/asadbekGo/market_system/models"
	"github.com/google/uuid"
)

type comingRepo struct {
	db *sql.DB
}

func NewComingRepo(db *sql.DB) *comingRepo {
	return &comingRepo{
		db: db,
	}
}

func (r *comingRepo) Create(req *models.CreateComing) (*models.Coming, error) {

	var (
		comingId = uuid.New().String()
		query    = `
			INSERT INTO "coming"(
				"id",
				"remaining_id",
				"updated_at"
			) VALUES ($1, $2, NOW())`
	)

	_, err := r.db.Exec(
		query,
		comingId,
		req.Remaining_Id,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(&models.ComingPrimaryKey{Id: comingId})
}

func (r *comingRepo) GetByID(req *models.ComingPrimaryKey) (*models.Coming, error) {

	var (
		coming models.Coming
		query  = `
			SELECT
			"id",
			"remaining_id",
			"created_at",
			"updated_at"	
			FROM "coming"
			WHERE "id" = $1
		`
	)

	err := r.db.QueryRow(query, req.Id).Scan(
		&coming.Id,
		&coming.Remaining_Id,
		&coming.CreatedAt,
		&coming.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &coming, nil
}

func (r *comingRepo) GetList(req *models.GetListComingRequest) (*models.GetListComingResponse, error) {
	var (
		resp   models.GetListComingResponse
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
		where += " AND 	remaining_id ILIKE" + " '%" + req.Search + "%'"
	}

	var query = `
		SELECT
			COUNT(*) OVER(),
			"id",
			"remaining_id", 
			"created_at",
			"updated_at"
		FROM "coming"
	`

	query += where + sort + offset + limit
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			coming models.Coming
		)

		err = rows.Scan(
			&resp.Count,
			&coming.Id,
			&coming.Remaining_Id,
			&coming.CreatedAt,
			&coming.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Comings = append(resp.Comings, &coming)
	}

	return &resp, nil
}

func (r *comingRepo) Update(req *models.UpdateComing) (int64, error) {

	query := `
		UPDATE coming
			SET
				"remaining_id"=$2, 
				updated_at=NOW()
		WHERE id = $1
	`
	result, err := r.db.Exec(
		query,
		req.Id,
		req.Remaining_Id,
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

func (r *comingRepo) Delete(req *models.ComingPrimaryKey) error {
	_, err := r.db.Exec("DELETE FROM coming WHERE id = $1", req.Id)
	return err
}
