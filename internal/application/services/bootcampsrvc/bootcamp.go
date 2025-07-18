package bootcampsrvc

import (
    "log"
	"context"
	"database/sql"
    // _"net/http"
	"github.com/melikadaryabak/bootcamp/internal/dto/entity"
)

type BootcampSrvc struct {
	db *sql.DB
}

func NewBootcampSrvc(db *sql.DB) BootcampSrvc {
	return BootcampSrvc{db: db}
}

func (b BootcampSrvc) GetBootcamps(ctx context.Context) ([]entity.Bootcamp, error) {
	bootcamps := []entity.Bootcamp{}

	// var bootcamps []Bootcamp

    // Get bootcamps query 
    rows, err := b.db.Query(`
        SELECT b.id, b.name, b.description, c.id, c.name
        FROM bootcamp b
        JOIN category c ON b.category_id = c.id
    `)

   // Handle database query error
   if err != nil {
    // http.Error(w, `{"error": "Failed to query bootcamps"}`, http.StatusInternalServerError)
    log.Println("Query error:", err)
    return nil ,err
}

defer rows.Close()
 // Scan rows into bootcamps
    for rows.Next() {
        var b entity.Bootcamp
        // var b Bootcamp
        if err := rows.Scan(
            &b.ID, &b.Name,
            &b.Description,
            &b.Category.ID,
            &b.Category.Name,
           ); err != nil {
            log.Println("Scan error:", err)
  
            // defer rows.Close()
            return nil,err
        }
        bootcamps = append(bootcamps, b)
    }

	return bootcamps,nil
}


func (b BootcampSrvc) PostBootcamp(ctx context.Context, bootcamp entity.Bootcamp) (int64, error) {

    // Post bootcamps query 
    query :=(`
    INSERT INTO bootcamp (name, description, category_id)
    VALUES (?, ?, ?)
    `)

   // Handle database query error
   result, err := b.db.ExecContext(ctx, query, bootcamp.Name, bootcamp.Description, bootcamp.Category.ID)
   if err != nil {
    log.Println("Query error:", err)
    return 0 ,err
}
insertedID, err := result.LastInsertId()
	if err != nil {
		log.Println("Failed to fetch last insert ID:", err)
		return 0 , err
	}

	return insertedID ,nil
}


func (b BootcampSrvc) DeleteBootcamp(ctx context.Context,id int64) (bool, error) {

    // Get bootcamps query 
    query := (`
    DELETE FROM bootcamp WHERE id = ?
    `)

   // Handle database query error
   result, err := b.db.ExecContext(ctx, query,id)
   if err != nil {
    log.Println("Query error:", err)
    return false ,err
}

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


func (b BootcampSrvc) PutBootcamp(ctx context.Context,bootcamp entity.Bootcamp) (bool, error) {

    // Get bootcamps query 
    query := (`
    UPDATE bootcamp
    SET name = ?, description = ?, category_id = ?
    WHERE id = ?
    `)

   // Handle database query error
   _ , err := b.db.ExecContext(ctx, query, bootcamp.Name, bootcamp.Description, bootcamp.Category.ID, bootcamp.ID)
   if err != nil {
    log.Println("Query error:", err)
    return false ,err
}
	return true,nil
}