package postgres

import (
	"database/sql"
	"fmt"

	"github.com/asadbekGo/market_system/models"
	"github.com/google/uuid"
)

type branchRepo struct {
	db *sql.DB
}

func NewBranchRepo(db *sql.DB) *branchRepo {
	return &branchRepo{
		db: db,
	}
}

func (r *branchRepo) Create(req *models.CreateBranch) (*models.Branch, error) {

	var (
		branchId = uuid.New().String()
		query    = `
			INSERT INTO "branch"(
				"id",
				"name", 
				"address",
				"updated_at"
			) VALUES ($1, $2,$3, NOW())`
	)

	_, err := r.db.Exec(
		query,
		branchId,
		req.Name,
		req.Address,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(&models.BranchPrimaryKey{Id: branchId})
}

func (r *branchRepo) GetByID(req *models.BranchPrimaryKey) (*models.Branch, error) {

	var (
		branch models.Branch
		query  = `
			SELECT
			"id",
			"name", 
			"address",
			"created_at",
			"updated_at"	
			FROM "branch"
			WHERE "id" = $1
		`
	)

	err := r.db.QueryRow(query, req.Id).Scan(
		&branch.Id,
		&branch.Name,
		&branch.Address,
		&branch.CreatedAt,
		&branch.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &branch, nil
}

func (r *branchRepo) GetList(req *models.GetListBranchRequest) (*models.GetListBranchResponse, error) {
	var (
		resp   models.GetListBranchResponse
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
			"address", 
			"created_at",
			"updated_at"
		FROM "branch"
	`

	query += where + sort + offset + limit
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			branch models.Branch
		)

		err = rows.Scan(
			&resp.Count,
			&branch.Id,
			&branch.Name,
			&branch.Address,
			&branch.CreatedAt,
			&branch.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Branchs = append(resp.Branchs, &branch)
	}

	return &resp, nil
}

func (r *branchRepo) Update(req *models.UpdateBranch) (int64, error) {

	query := `
		UPDATE branch
			SET
				name=$2, 
				address=$3, 
				updated_at=NOW()
		WHERE id = $1
	`
	result, err := r.db.Exec(
		query,
		req.Id,
		req.Name,
		req.Address,
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

func (r *branchRepo) Delete(req *models.BranchPrimaryKey) error {
	_, err := r.db.Exec("DELETE FROM branch WHERE id = $1", req.Id)
	return err
}
