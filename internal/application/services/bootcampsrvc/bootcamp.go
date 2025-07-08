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
            &b.Category.Name
           ); err != nil {
            log.Println("Scan error:", err)
  
            // defer rows.Close()
            return nil,err
        }
        bootcamps = append(bootcamps, b)
    }

	return bootcamps
}