package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Message   string `json:"message"`
	Version   string `json:"version"`
	Language  string `json:"language"`
	Timestamp string `json:"timestamp"`
}

type HealthResponse struct {
	Status    string `json:"status"`
	Language  string `json:"language"`
	Framework string `json:"framework"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Message:   "Hello from Go!",
		Version:   "1.0.0",
		Language:  "Go",
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HealthResponse{
		Status:    "healthy",
		Language:  "Go",
		Framework: "net/http",
	})
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := []User{
		{ID: 1, Name: "Alice", City: "Mumbai"},
		{ID: 2, Name: "Bob", City: "Delhi"},
		{ID: 3, Name: "Carol", City: "Pune"},
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/users", usersHandler)

	fmt.Printf("Go server running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
