package services

import (
	// "database/sql"
	"github.com/melikadaryabak/bootcamp/internal/application/services/categorysrvc"
	"github.com/melikadaryabak/bootcamp/internal/application/services/bootcampsrvc"
	"github.com/melikadaryabak/bootcamp/internal/infrastructure/db"
)

type Services struct {
	CategorySrvc categorysrvc.CategorySrvc
	BootcampSrvc bootcampsrvc.BootcampSrvc
}

func NewServices(repo *db.BootcampRepo) *Services {
	return &Services{
		// CategorySrvc: categorysrvc.NewCategorySrvc(repo),
		BootcampSrvc: bootcampsrvc.NewBootcampSrvc(repo),
	}
}
