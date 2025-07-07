package bootcampsrvc

import (
	"context"
	"database/sql"
	"github.com/melikadaryabak/bootcamp/internal/dto/entity"
	"database/sql"
)

type BootcampSrvc struct {
	db *sql.DB
}

func NewBootcampSrvc(db *sql.DB) BootcampSrvc {
	return BootcampSrvc{db: db}
}

func (b BootcampSrvc) GetBootcamps(ctx context.Context) []entity.Bootcamp {
	bootcamps := []entity.Bootcamp{}

	// var bootcamps []Bootcamp

    // Get bootcamps query 
    rows, err := db.Query(`
        SELECT b.id, b.name, b.description, c.id, c.name
        FROM bootcamp b
        JOIN category c ON b.category_id = c.id
    `)

   // Handle database query error
   if err != nil {
    http.Error(w, `{"error": "Failed to query bootcamps"}`, http.StatusInternalServerError)
    log.Println("Query error:", err)
    return
}
    
 // Scan rows into bootcamps
    for rows.Next() {
        var b Bootcamp
        if err := rows.Scan(&b.ID, &b.Name, &b.Description, &b.Category.ID, &b.Category.Name); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            defer rows.Close()
            return
        }
        bootcamps = append(bootcamps, b)
    }

	return bootcamps
}