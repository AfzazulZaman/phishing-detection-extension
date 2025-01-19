package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PhishingReport struct {
	URL string `json:"url"`
}

func main() {
	http.HandleFunc("/report", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var report PhishingReport
			err := json.NewDecoder(r.Body).Decode(&report)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}
			fmt.Println("Reported phishing URL:", report.URL)
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
