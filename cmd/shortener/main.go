package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var store = map[string]string{}

const uuid = "i2o3hgo3ihg"

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost && r.URL.Path == "/" {
		url, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read URL", http.StatusBadRequest)
			return
		}
		store[uuid] = string(url)
		fmt.Fprintf(w, "https://localhost:8080/%s", uuid)
	}
	if r.Method == http.MethodGet {
		if r.URL.Path == "/" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}
		if r.URL.Path != "/" {
			id := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/"))

			if id == "" {
				http.Error(w, "ID is empty", http.StatusBadRequest)
				return
			}

			url, ok := store[uuid]
			if !ok || strings.TrimSpace(url) == "" {
				http.Error(w, "URL not found", http.StatusNotFound)
				return
			}

			if id != uuid {
				http.Error(w, "Invalid ID", http.StatusNotFound)
				return
			}

			http.Redirect(w, r, store[uuid], http.StatusFound)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", redirectHandler)

	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
