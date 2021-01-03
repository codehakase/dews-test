package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type user struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {

	port := os.Getenv("PORT")

	http.HandleFunc("/", healthCheckHandler)
	http.HandleFunc("/users/create", createUser)
	http.HandleFunc("/users/list", listUsers)

	log.Println("Running web server on PORT", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var data user
	// decode request
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("failed to decode request data: %v", err)
		ex(w, "failed to decode request", http.StatusBadRequest)
	}

	// persit data
	err = WriteFile(data)
	if err != nil {
		log.Printf("failed to save user data: %v", err)
		ex(w, "failed to save user data", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "User created",
	})
}

func ex(w http.ResponseWriter, message string, code int) {
	http.Error(w, message, code)
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := ReadUserFile()
	if err != nil {
		ex(w, "error retrieving users", http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}
