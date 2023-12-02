package postgres

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"app/model"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) Create(req model.UserCreate) (*model.User, error) {
	var userId = uuid.New().String()

	var user = model.User{
		Id:   userId,
		Name: req.Name,
	}

	var query = `INSERT INTO users(id, name, updated_at) VALUES($1, $2, NOW())`

	_, err := u.db.Exec(query, user.Id, user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetById(req model.UserPrimaryKey) (*model.User, error) {

	var user = model.User{}

	var query = `SELECT 
			id, 
			name, 
			created_at, 
			updated_at
			FROM 
				users
			WHERE id = $1
	`

	resp := u.db.QueryRow(query, req.Id)

	if resp.Err() != nil {
		return nil, resp.Err()
	}
	err := resp.Scan(
		&user.Id,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetList(req model.UserGetListRequest) (*model.UserGetListResponse, error) {

	var resp = model.UserGetListResponse{}
	var offset string = " offset"
	var limit string = " limit"

	if req.Offset <= 0 {
		offset += " 0"
	} else if req.Offset > 0 {
		offset += fmt.Sprintf(" %d", req.Offset)
	}

	if req.Limit <= 0 {
		limit += " 10"
	} else if req.Limit > 0 {
		limit += fmt.Sprintf(" %d", req.Limit)
	}

	var query = `SELECT 
			COUNT(*) OVER(),
			id, 
			name, 
			created_at, 
			updated_at
			FROM 
				users
	`
	query += offset + limit
	fmt.Println(query)
	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user = model.User{}
		rows.Scan(
			&resp.Count,
			&user.Id,
			&user.Name,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		resp.Users = append(resp.Users, user)
	}

	rows.Close()

	return &resp, nil
}

func (u *UserRepo) Update(req model.UserUpdate) (int64, error) {
	var (
		query = `
				UPDATE users SET name=$2,updated_at=NOW() WHERE id = $1
			`
	)

	resp, err := u.db.Exec(query, req.Id, req.Name)

	if err != nil {
		return resp.RowsAffected()
	}

	return resp.RowsAffected()

}

func (u *UserRepo) Delete(req model.UserPrimaryKey) (int64, error) {
	var (
		query = `
				DELETE FROM users WHERE id = $1
			`
	)

	resp, err := u.db.Exec(query, req.Id)

	if err != nil {
		return resp.RowsAffected()
	}

	return resp.RowsAffected()

}
