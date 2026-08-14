package main

import (
	"log"
	"net/http"

	"product-catalog/backend/handlers"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/products" {
		if r.Method != http.MethodGet {
			http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
			return
		}

		handlers.GetProducts(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	handlers.GetProductByID(w, r)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/products", productsHandler)
	mux.HandleFunc("/api/products/", productsHandler)

	handler := enableCORS(mux)

	log.Println("Server running on http://localhost:8000")

	err := http.ListenAndServe(":8000", handler)

	if err != nil {
		log.Fatal(err)
	}
}