package main

import (
    "encoding/json"
    "log"
    "net/http"
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