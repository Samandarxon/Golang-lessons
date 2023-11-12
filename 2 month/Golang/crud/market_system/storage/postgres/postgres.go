package postgres

import (
	"database/sql"
	"fmt"

	"github.com/asadbekGo/market_system/config"
	"github.com/asadbekGo/market_system/storage"

	_ "github.com/lib/pq"
)

type Store struct {
	db        *sql.DB
	category  storage.CategoryRepoI
	product   storage.ProductRepoI
	remaining storage.RemainingRepoI
	branch    storage.BranchRepoI
	coming    storage.ComingRepoI
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

func (s *Store) Category() storage.CategoryRepoI {

	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}

	return s.category
}

func (s *Store) Product() storage.ProductRepoI {

	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}

	return s.product
}
func (s *Store) Remaining() storage.RemainingRepoI {

	if s.remaining == nil {
		s.remaining = NewRemainingRepo(s.db)
	}

	return s.remaining
}
func (s *Store) Branch() storage.BranchRepoI {

	if s.branch == nil {
		s.branch = NewBranchRepo(s.db)
	}

	return s.branch
}

func (s *Store) Coming() storage.ComingRepoI {

	if s.coming == nil {
		s.coming = NewComingRepo(s.db)
	}

	return s.coming
}
