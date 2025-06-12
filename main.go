package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SumRequest struct {
	Numbers []float64 `json:"numbers"`
}

type SumResponse struct {
	Sum float64 `json:"sum"`
}

func sum(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var data SumRequest
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	sum := 0.0
	for _, number := range data.Numbers {
		sum += number
	}

	response := SumResponse{Sum: sum}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/sum", sum)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
