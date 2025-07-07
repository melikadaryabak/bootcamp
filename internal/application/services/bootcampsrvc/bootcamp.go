package bootcampsrvc

import (
	"context"
	"database/sql"
	"github.com/melikadaryabak/bootcamp/internal/dto/entity"
)

type BootcampSrvc struct {
	db *sql.DB
}

func NewBootcampSrvc(db *sql.DB) BootcampSrvc {
	return BootcampSrvc{db: db}
}

// func (b BootcampSrvc) GetCategories(ctx context.Context) []entity.Category {
// 	categories := []entity.Category{}

// 	//TODO: get categoies from sql
// 	return categories
// }