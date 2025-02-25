package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Status  string `json:"status"`
	Time    string `json:"time"`
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// set the response header to indicate you are to use json response
	w.Header().Set("Content-Type", "application/json")

	res := Response{
		Status:  "success: api running",
		Time:    time.Now().Format(time.RFC3339),
		Message: "CI/CD pipeline with Jenkins, Docker, Helm, and Kubernetes,",
	}

	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("starting server on port 8080...")

	http.ListenAndServe(":8080", nil)
}
