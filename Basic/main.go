// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Numbers struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Result struct {
	Sum int `json:"sum"`
}

func addNumbersHandler(w http.ResponseWriter, r *http.Request) {
	var numbers Numbers
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := Result{Sum: numbers.A + numbers.B}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func mains() {
	http.HandleFunc("/add", addNumbersHandler)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
