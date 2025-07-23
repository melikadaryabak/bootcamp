package db

import (
	"log"
	"context"
	"database/sql"
	"github.com/melikadaryabak/bootcamp/internal/dto/entity"
)

type BootcampRepo struct {
	db *sql.DB
}

func NewBootcampRepo(db *sql.DB) *BootcampRepo {
	return &BootcampRepo{db: db}
}

func (r *BootcampRepo) GetBootcamps(ctx context.Context) ([]entity.Bootcamp, error) {
	var bootcamps []entity.Bootcamp

	//Get bootcamps query 
    rows, err := r.db.QueryContext(ctx,`
        SELECT b.id, b.name, b.description, c.id, c.name
        FROM bootcamp b
        JOIN category c ON b.category_id = c.id
    `)
	if err != nil {
		return nil, err
	}

	//Scan rows into bootcamps
	defer rows.Close()
    for rows.Next() {
        var b entity.Bootcamp
        if err := rows.Scan(&b.ID, &b.Name, &b.Description, &b.Category.ID, &b.Category.Name); err != nil {
			return nil, err
        }
        bootcamps = append(bootcamps, b)
    }
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return bootcamps, nil
}


func (r *BootcampRepo) PostBootcamps(ctx context.Context, b entity.Bootcamp) (int64, error) {
	// var bootcamps []entity.Bootcamp

	// Post bootcamps query 
    result, err := r.db.ExecContext(ctx,`
    INSERT INTO bootcamp (name, description, category_id)
    VALUES (?, ?, ?)
	`, b.Name, b.Description, b.Category.ID)
	if err != nil {
		return 0, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		log.Println("Failed to fetch last insert ID:", err)
		return 0 , err
	}

	return insertedID ,nil
}


func (r *BootcampRepo) DeleteBootcamps(ctx context.Context,id int64) (bool, error) {

	// Get bootcamps query 
	result, err := r.db.ExecContext(ctx,`
    DELETE FROM bootcamp WHERE id = ?
	`,id)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Failed to get affected rows:", err)
		return false, err
	}

	if rowsAffected == 0 {
        log.Printf("No bootcamp found with id: %d", id)
		return false, nil
	}

	return true,nil
}


func (r *BootcampRepo) PutBootcamps(ctx context.Context,bootcamp entity.Bootcamp) (bool, error) {

	// Get bootcamps query 
    _ , err := r.db.ExecContext(ctx,`
    UPDATE bootcamp
    SET name = ?, description = ?, category_id = ?
    WHERE id = ?
    `, bootcamp.Name, bootcamp.Description, bootcamp.Category.ID, bootcamp.ID)

	if err != nil {
		log.Println("Failed to execute update:", err)
		return false, err
	}
	return true, nil
}