package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"product-catalog/backend/models"
)

func loadProducts() ([]models.Product, error) {
	data, err := os.ReadFile("data/products.json")
	if err != nil {
		return nil, err
	}

	var products []models.Product

	err = json.Unmarshal(data, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := loadProducts()

	if err != nil {
		http.Error(w, `{"error":"Failed to load products"}`, http.StatusInternalServerError)
		return
	}

	category := r.URL.Query().Get("category")

	if category != "" {
		var filteredProducts []models.Product

		for _, product := range products {
			if strings.EqualFold(product.Category, category) {
				filteredProducts = append(filteredProducts, product)
			}
		}

		products = filteredProducts
	}

	json.NewEncoder(w).Encode(products)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := loadProducts()

	if err != nil {
		http.Error(w, `{"error":"Failed to load products"}`, http.StatusInternalServerError)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/api/products/")

	for _, product := range products {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	http.Error(w, `{"error":"Product not found"}`, http.StatusNotFound)
}
