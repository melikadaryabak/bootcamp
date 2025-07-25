package categorysrvc

import (
	"context"
	"github.com/melikadaryabak/bootcamp/internal/dto/entity"
    "github.com/melikadaryabak/bootcamp/internal/infrastructure/db"
)

type CategorySrvc struct {
	repo *db.CategoryRepo
}

func NewCategorySrvc(repo *db.CategoryRepo) *CategorySrvc {
	return &CategorySrvc{repo: repo}
}

func (c CategorySrvc) GetCategories(ctx context.Context) ([]entity.Category,error) {
    return c.repo.GetCategories(ctx)
}
