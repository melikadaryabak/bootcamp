package db

import (
	"log"
	"context"
	"database/sql"
	"github.com/melikadaryabak/bootcamp/internal/dto/entity"
)

type CategoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (r *CategoryRepo) GetCategories(ctx context.Context) ([]entity.Category, error) {
	var categories []entity.Category

	//Get categories query 
    rows, err := r.db.QueryContext(ctx,`
	SELECT c.id, c.name
	FROM category c
    `)
if err != nil {
	log.Println("Failed to query categories:", err)
	return nil, err
}

	//Scan rows into categories
	defer rows.Close()
    for rows.Next() {
        var c entity.Category
        if err := rows.Scan(&c.ID,&c.Name,); err != nil {
			return nil, err
        }
        categories = append(categories, c)
    }
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}
