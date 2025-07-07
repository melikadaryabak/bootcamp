package categorysrvc

import (
	"context"
	"database/sql"
	"github.com/melikadaryabak/bootcamp/internal/dto/entity"
)

type CategorySrvc struct {
	db *sql.DB
}

func NewCategorySrvc(db *sql.DB) CategorySrvc {
	return CategorySrvc{db: db}
}

func (c CategorySrvc) GetCategories(ctx context.Context) []entity.Category {
	categories := []entity.Category{}

	//TODO: get categoies from sql
	return categories
}
