package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/melikadaryabak/bootcamp/internal/application/services"
	"github.com/melikadaryabak/bootcamp/internal/dto/entity"
	_ "github.com/go-sql-driver/mysql"
)


type Server struct {
	srvc services.Services
}


func NewServer(port string, srvc services.Services) error {

	server := Server{
		srvc: srvc,
	}

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			server.GetCategories(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/bootcamps", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			server.GetBootcamps(w, r)
		case http.MethodPost:
			server.PostBootcamp(w, r)
		case http.MethodDelete :
			server.DeleteBootcamp(w, r)
		case http.MethodPut :
			server.PutBootcamp(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
	return nil
}


func (s Server) GetCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// categories , err := s.srvc.CategorySrvc.GetCategories(r.Context())
	// if err != nil{
	//   http.Error(w, fmt.Sprintf("error for get Categories: %w" , err), http.StatusInternalServerError)
	// }

	categories , err := s.srvc.CategorySrvc.GetCategories(r.Context())
	if err != nil {
		log.Printf("error getting categories: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	// Set JSON header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode bootcamps to JSON
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"categories": categories}); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}


func (s Server) GetBootcamps(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// bootcamps , err := s.srvc.BootcampSrvc.GetBootcamps(r.Context())
	// if err != nil{
	//   http.Error(w, fmt.Sprintf("error for get bootcamps: %w" , err), http.StatusInternalServerError)
	//   return
	// }

	bootcamps , err := s.srvc.BootcampSrvc.GetBootcamps(r.Context())
	if err != nil {
		log.Printf("error getting bootcamps: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set JSON header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode bootcamps to JSON
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"bootcamps": bootcamps}); err != nil {
        log.Printf("Error encoding response: %v", err)
    }
}


func (s Server) PostBootcamp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Decode request body into Bootcamp struct
    var bootcamp entity.Bootcamp
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&bootcamp); err != nil {
        http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
        return
    }
	defer r.Body.Close()

	// Check that name is not empty
if bootcamp.Name == "" {
    http.Error(w, `{"error": "Name is required"}`, http.StatusBadRequest)
    return
}

	// bootcampId , err := s.srvc.BootcampSrvc.PostBootcamp(r.Context(),bootcamp)
	// if err != nil{
	//   http.Error(w, fmt.Sprintf("error for add bootcamps: %w" , err), http.StatusInternalServerError)
	//   return
	// }

	insertedID , err := s.srvc.BootcampSrvc.PostBootcamp(r.Context(),bootcamp)
	if err != nil {
		log.Printf("error for add bootcamp: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set JSON header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Encode bootcamps to JSON
	if err := json.NewEncoder(w).Encode(insertedID); err != nil {
        log.Printf("Error encoding response: %v", err)
    }
}


func (s Server) DeleteBootcamp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	idInt64 := int64(id)

	// success, err := s.srvc.BootcampSrvc.DeleteBootcamp(r.Context(), idInt64)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Error deleting bootcamp: %v", err), http.StatusInternalServerError)
	// 	return
	// }

	success , err := s.srvc.BootcampSrvc.DeleteBootcamp(r.Context(),idInt64)
	if err != nil {
		log.Printf("error deleting bootcamp: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if !success {
		http.Error(w, "Bootcamp not found", http.StatusNotFound)
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


func (s Server) PutBootcamp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Decode request body into Bootcamp struct
    var bootcamp entity.Bootcamp
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&bootcamp); err != nil {
        http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
        return
    }
	defer r.Body.Close()

	// Check that name is not empty
if bootcamp.Name == "" {
    http.Error(w, `{"error": "Name is required"}`, http.StatusBadRequest)
    return
}

	if bootcamp.ID == 0 {
		http.Error(w, "Missing bootcamp ID", http.StatusBadRequest)
		return
	}

	// success, err := s.srvc.BootcampSrvc.PutBootcamp(r.Context(), bootcamp)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Error updating bootcamp: %v", err), http.StatusInternalServerError)
	// 	return
	// }

	success , err := s.srvc.BootcampSrvc.PutBootcamp(r.Context(),bootcamp)
	if err != nil {
		log.Printf("error updating bootcamp: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if !success {
		http.Error(w, "Bootcamp not found", http.StatusNotFound)
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
