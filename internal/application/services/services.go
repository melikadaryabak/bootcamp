package services

import (
	"database/sql"
	"github.com/melikadaryabak/bootcamp/internal/application/services/categorysrvc"
	"github.com/melikadaryabak/bootcamp/internal/application/services/bootcampsrvc"
)

type Services struct {
	CategorySrvc categorysrvc.CategorySrvc
	BootcampSrvc bootcampsrvc.BootcampSrvc
}

func NewServices(db *sql.DB) Services {
	return Services{
		CategorySrvc: categorysrvc.NewCategorySrvc(db),
		BootcampSrvc: bootcampsrvc.NewBootcampSrvc(db),
	}
}
