package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"app/config"
	"app/storage"
)

type Store struct {
	db   *sql.DB
	user *UserRepo
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {
	connect := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	)

	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}

	return &Store{
		db: db,
	}, nil
}

func (s *Store) User() storage.UserRepoI {
	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}
	return s.user
}
