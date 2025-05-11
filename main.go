package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Greeting struct {
    Message string `json:"message"`
}

func main() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.HandleFunc("/api/greet", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
        greeting := Greeting{Message: "Hello, Junior DevSecOps!"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(greeting)
    })
    fmt.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
} 
