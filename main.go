package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
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

var bootcamps = []Bootcamp{
	{
        ID: 1,
        Name: "Go Bootcamp",   
        Description: "This bootcamp teaches the fundamentals of Go programming, covering goroutines, channels, and building web APIs.",
        Category: Categories[0],
        },

	{
        ID: 2,
        Name: "PHP Bootcamp",
        Description: "A complete PHP development course focusing on server-side scripting, working with databases, and building dynamic websites.",
        Category: Categories[0],
        },
        {
            ID:          3,
            Name:        "Android Bootcamp",
            Description: "Learn to build modern Android apps using Java and Kotlin, including UI design, data storage, and integration with RESTful APIs.",
            Category:    Categories[2],
        },
        {
            ID:         4,
            Name:        "c++ Bootcamp",
            Description: "Learn to build modern Android apps using Java and Kotlin, including UI design, data storage, and integration with RESTful APIs.",
            Category:    Categories[1],
        },
}

func main() {
    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/bootcamps", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
        case http.MethodGet :
                bootcampsHandler(w, r)
            case http.MethodPost :
                newbootcampsHandler(w, r)
            case http.MethodDelete :
                deletebootcampsHandler(w, r)
            default:
                http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        }
	})

    // Start the server on port 8080
    log.Println("Server listening on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // Only allow GET requests
    if r.Method != http.MethodGet {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    // Prepare the response payload
    payload := map[string]string{"message": "Hello, World!"}

    // Set JSON header
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    // Encode payload to JSON
    if err := json.NewEncoder(w).Encode(payload); err != nil {
        log.Printf("Error encoding response: %v", err)
    }
}

func bootcampsHandler(w http.ResponseWriter, r *http.Request) {
    // Only allow GET requests
    if r.Method != http.MethodGet {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    // Prepare the response payload
    payload := map[string]interface{}{
        "bootcamps": bootcamps,
    }

    // Set JSON header
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    // Encode payload to JSON
    if err := json.NewEncoder(w).Encode(payload); err != nil {
        log.Printf("Error encoding response: %v", err)
    }
}
func newbootcampsHandler(w http.ResponseWriter, r *http.Request) {

    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    var newBootcamp Bootcamp
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&newBootcamp); err != nil {
        http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
        return
    }
    
    newBootcamp.ID = len(bootcamps) + 1
    bootcamps = append(bootcamps, newBootcamp)

    // Set JSON header
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)


    // Encode payload to JSON
    if err := json.NewEncoder(w).Encode(newBootcamp); err != nil {
        log.Printf("Error encoding response: %v", err)
    }
}
func deletebootcampsHandler(w http.ResponseWriter, r *http.Request) {
    
    if r.Method != http.MethodDelete {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
        
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
    
        
        found := false
        for i, b := range bootcamps {
            if b.ID == id {
                bootcamps = append(bootcamps[:i], bootcamps[i+1:]...)
                found = true
                break
            }
        }
    
        if !found {
            http.Error(w, `{"error": "Bootcamp not found"}`, http.StatusNotFound)
            return
        }
    
        // Set JSON header
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)

        if err := json.NewEncoder(w).Encode(map[string]string{"message": "Bootcamp deleted successfully"}); err != nil {
            log.Printf("Error encoding response: %v", err)
        }
        
    }
    
