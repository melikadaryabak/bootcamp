package db

import (
	"database/sql"
)

type Repository struct {
	BootcampRepo *BootcampRepo
	CategoryRepo *CategoryRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		BootcampRepo: NewBootcampRepo(db),
		CategoryRepo: NewCategoryRepo(db),
	}
}