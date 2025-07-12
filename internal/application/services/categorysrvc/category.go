package categorysrvc

import (
	"log"
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

func (c CategorySrvc) GetCategories(ctx context.Context) ([]entity.Category,error) {
	categories := []entity.Category{}

	 // Get category query 
	 rows, err := c.db.Query(`
	 SELECT c.id, c.name
	 FROM category c
 `)

 // Handle database query error
 if err != nil {
    log.Println("Query error:", err)
    return nil ,err
}

defer rows.Close()
 // Scan rows into bootcamps
    for rows.Next() {
        var c entity.Category
        if err := rows.Scan(
            &c.ID,
            &c.Name,
           ); err != nil {
            log.Println("Scan error:", err)
  
            // defer rows.Close()
            return nil,err
        }
        categories = append(categories, c)
    }
	return categories, nil
}
