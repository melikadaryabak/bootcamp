package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var (
    db *sql.DB
 err error
) 

type Category struct{
    ID int `json:"id"`
    Name string `json:"name"`
}

var Categories = []Category{
    {ID: 1, Name:"web"},
    {ID: 2, Name:"windows"},
    {ID: 3, Name: "android"},
}

type Bootcamp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
    Description string `json:"description"`
    Category Category `json:"category"`
}

func main() {

    // Connect to MySQL
    db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/amacodecamp")

    // Check database connection
    if err != nil {
        log.Fatalf("Error opening database connection: %v", err)
    }
    if err = db.Ping(); err != nil {
        log.Fatalf("Error pinging database: %v", err)
    }
    
   // Static file handler
    http.Handle("/", http.FileServer(http.Dir("./static")))

    // Handle bootcamps API
    http.HandleFunc("/bootcamps", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
        case http.MethodGet :
                bootcampsHandler(w, r)
            case http.MethodPost :
                newbootcampsHandler(w, r)
            case http.MethodDelete :
                deletebootcampsHandler(w, r)
            case http.MethodPut :
                editbootcampsHandler(w, r)
            default:
                http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        }
	})

    // Start HTTP server on port 8080
    log.Println("Server listening on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}


func bootcampsHandler(w http.ResponseWriter, r *http.Request) {

    // Only allow GET requests
    if r.Method != http.MethodGet {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    var bootcamps []Bootcamp

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
 defer rows.Close()
    for rows.Next() {
        var b Bootcamp
        if err := rows.Scan(&b.ID, &b.Name, &b.Description, &b.Category.ID, &b.Category.Name); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
           
            return
        }
        bootcamps = append(bootcamps, b)
    }

    // Set JSON header
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    // Encode bootcamps to JSON
    if err := json.NewEncoder(w).Encode(map[string]interface{}{"bootcamps": bootcamps}); err != nil {
        log.Printf("Error encoding response: %v", err)
    }
}


func newbootcampsHandler(w http.ResponseWriter, r *http.Request) {

 // Only allow Post requests
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    // Decode request body into Bootcamp struct
    var newBootcamp Bootcamp
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&newBootcamp); err != nil {
        http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
        return
    }

    // Check for missing category ID
    if newBootcamp.Category.ID == 0 {
        http.Error(w, `{"error": "Missing category ID"}`, http.StatusBadRequest)
        return
    }

// Check that name is not empty
if newBootcamp.Name == "" {
    http.Error(w, `{"error": "Name is required"}`, http.StatusBadRequest)
    return
}

    // Insert new bootcamp into database
    result, err := db.Exec(
        "INSERT INTO bootcamp (name, description, category_id) VALUES (?, ?, ?)",
        newBootcamp.Name, newBootcamp.Description, newBootcamp.Category.ID,
    )
    
    // Handle database insert error
    if err != nil {
        http.Error(w, `{"error": "Failed to insert into database"}`, http.StatusInternalServerError)
        log.Println("Insert error:", err)
        return
    }
    
    // Get ID of newly inserted bootcamp
    insertedID, err := result.LastInsertId()
    if err != nil {
        http.Error(w, `{"error": "Failed to retrieve inserted ID"}`, http.StatusInternalServerError)
        return
    }

    // Set ID on newBootcamp struct
    newBootcamp.ID = int(insertedID)
   
    // Set JSON header
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    // Encode new bootcamp to JSON
    if err := json.NewEncoder(w).Encode(newBootcamp); err != nil {
        log.Printf("Error encoding response: %v", err)
    }
}


func deletebootcampsHandler(w http.ResponseWriter, r *http.Request) {
    
    // Only allow Delete requests
    if r.Method != http.MethodDelete {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
        // Extract 'id' from URL query parameters
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, `{"error": "ID is required"}`, http.StatusBadRequest)
        return
    }
    
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, `{"error": "Invalid ID"}`, http.StatusBadRequest)
        return
    }
    
    // Delete bootcamp by id  
    result, err := db.Exec("DELETE FROM bootcamp WHERE id = ?", id)
    if err != nil {
        http.Error(w, `{"error": "Database error"}`, http.StatusInternalServerError)
        log.Printf("Delete error: %v", err)
        return
    }
    
    // Check delete result and handle not found bootcamp
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        http.Error(w, `{"error": "Could not determine result"}`, http.StatusInternalServerError)
        return
    }
    if rowsAffected == 0 {
        http.Error(w, `{"error": "Bootcamp not found"}`, http.StatusNotFound)
        return
    }
    
        // Set JSON header
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)

       // Send success response after deletion
        if err := json.NewEncoder(w).Encode(map[string]string{"message": "Bootcamp deleted successfully"}); err != nil {
            log.Printf("Error encoding response: %v", err)
        }
        
    }


func editbootcampsHandler(w http.ResponseWriter, r *http.Request) {
    
            // Only allow Put requests
            if r.Method != http.MethodPut {
                http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
                return
            }
                
              // Extract 'id' from URL query parameters
                idStr := r.URL.Query().Get("id")
                if idStr == "" {
                    http.Error(w, `{"error": "ID is required"}`, http.StatusBadRequest)
                    return
                }
            
                // Convert 'id' from string to int
                id, err := strconv.Atoi(idStr)
                if err != nil {
                    http.Error(w, `{"error": "Invalid ID"}`, http.StatusBadRequest)
                    return
                }
    
      // Decode request body into Bootcamp struct
       var update Bootcamp
        decoder := json.NewDecoder(r.Body)
        if err := decoder.Decode(&update); err != nil {
            http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
            return
        }

        // Check that name is not empty
if update.Name == "" {
    http.Error(w, `{"error": "Name is required"}`, http.StatusBadRequest)
    return
}
    
      // update bootcamp in the database
    _ , err = db.Exec(`UPDATE bootcamp SET name = ?, description = ?, category_id = ? WHERE id = ?`,
    update.Name, update.Description, update.Category.ID, id)
if err != nil {
    http.Error(w, `{"error": "Database error"}`, http.StatusInternalServerError)
    log.Printf("DB update error: %v", err)
    return
}
    
        // Set JSON header
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
    
        //Encode updated bootcamp to JSON
        if err := json.NewEncoder(w).Encode(map[string]string{"message": "Bootcamp edit successfully"}); err != nil {
            log.Printf("Error encoding response: %v", err)
        }
      }