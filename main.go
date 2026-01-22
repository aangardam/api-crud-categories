package main

import (
	"api-crud-categories/handlers"
	"api-crud-categories/helpers"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API is running",
		})
	})

	http.HandleFunc("/api/v1/categories", handlers.CategoryHandler)
	http.HandleFunc("/api/v1/categories/", handlers.CategoryByIDHandler)

	port, err := helpers.FindAvaliblePort(3000)
	if err != nil {
		fmt.Println("Failed to find avalible port", port)
		return
	}

	fmt.Println("Server running on port", port)
	err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println("Failed to run server")
	}
}
