package db

import (
	"context"
	"database/sql"
	"github.com/melikadaryabak/bootcamp/internal/dto/entity"
)

type BootcampRepo struct {
	DB *sql.DB
}

func NewBootcampRepo(db *sql.DB) *BootcampRepo {
	return &BootcampRepo{DB: db}
}

func (r *BootcampRepo) GetBootcamps(ctx context.Context) ([]entity.Bootcamp, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT b.id, b.name, b.description, c.id, c.name
		FROM bootcamp b
		JOIN category c ON b.category_id = c.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bootcamps []entity.Bootcamp
	for rows.Next() {
		var b entity.Bootcamp
		err := rows.Scan(&b.ID, &b.Name, &b.Description, &b.Category.ID, &b.Category.Name)
		if err != nil {
			return nil, err
		}
		bootcamps = append(bootcamps, b)
	}
	return bootcamps, nil
}

// var bootcamps []Bootcamp

//     // Get bootcamps query 
//     rows, err := db.Query(`
//         SELECT b.id, b.name, b.description, c.id, c.name
//         FROM bootcamp b
//         JOIN category c ON b.category_id = c.id
//     `)
    
//  // Scan rows into bootcamps
//     for rows.Next() {
//         var b Bootcamp
//         if err := rows.Scan(&b.ID, &b.Name, &b.Description, &b.Category.ID, &b.Category.Name); err != nil {
//             http.Error(w, err.Error(), http.StatusInternalServerError)
//             defer rows.Close()
//             return
//         }
//         bootcamps = append(bootcamps, b)
//     }
