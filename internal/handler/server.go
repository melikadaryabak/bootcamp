package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"github.com/melikadaryabak/bootcamp/internal/application/services"
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

	categories , err := s.srvc.CategorySrvc.GetCategories(r.Context())
	if err != nil{
	  http.Error(w, fmt.Sprintf("error for get Categories: %w" , err), http.StatusInternalServerError)
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
	
	bootcamps , err := s.srvc.BootcampSrvc.GetBootcamps(r.Context())
	if err != nil{
	  http.Error(w, fmt.Sprintf("error for get bootcamps: %w" , err), http.StatusInternalServerError)
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

	bootcampId , err := s.srvc.BootcampSrvc.PostBootcamp(r.Context(),bootcamp)
	if err != nil{
	  http.Error(w, fmt.Sprintf("error fo add bootcamps: %w" , err), http.StatusInternalServerError)
	}

	// Set JSON header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode bootcamps to JSON
	if err := json.NewEncoder(w).Encode(bootcamp); err != nil {
        log.Printf("Error encoding response: %v", err)
    }
}
