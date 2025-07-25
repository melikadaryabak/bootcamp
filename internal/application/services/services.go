package services

import (
	// "database/sql"
	"github.com/melikadaryabak/bootcamp/internal/application/services/categorysrvc"
	"github.com/melikadaryabak/bootcamp/internal/application/services/bootcampsrvc"
	repopkg "github.com/melikadaryabak/bootcamp/internal/infrastructure/db"
)

type Services struct {
	CategorySrvc *categorysrvc.CategorySrvc
	BootcampSrvc *bootcampsrvc.BootcampSrvc
}

func NewServices(repo *repopkg.Repository) *Services {
	return &Services{
		CategorySrvc: categorysrvc.NewCategorySrvc(repo.CategoryRepo),
		BootcampSrvc: bootcampsrvc.NewBootcampSrvc(repo.BootcampRepo),
	}
}
