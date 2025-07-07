package services

import (
	"database/sql"

	"github.com/melikadaryabak/bootcamp/internal/application/services/categorysrvc"
)

type Services struct {
	CategorySrvc categorysrvc.CategorySrvc
}

func NewServices(db *sql.DB) Services {
	return Services{
		CategorySrvc: categorysrvc.NewCategorySrvc(db),
	}
}
