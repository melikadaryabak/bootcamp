package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Bootcamp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var bootcamps = []Bootcamp{
	{ID: 1, Name: "Go Bootcamp"},
	{ID: 2, Name: "PHP Bootcamp"},
}

func main() {
    //1
    http.HandleFunc("/hello", helloHandler)
    //2
    http.HandleFunc("/bootcamps", bootcampsHandler)
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