package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

var products = []Product{
	{
		ID:       "1",
		Name:     "Cotton Kurta",
		Price:    1299,
		Category: "Clothing",
	},
	{
		ID:       "2",
		Name:     "Steel Water Bottle",
		Price:    649,
		Category: "Kitchen",
	},
	{
		ID:       "3",
		Name:     "Leather Wallet",
		Price:    899,
		Category: "Accessories",
	},
	{
		ID:       "4",
		Name:     "Ceramic Mug",
		Price:    399,
		Category: "Kitchen",
	},
	{
		ID:       "5",
		Name:     "Running Shoes",
		Price:    2499,
		Category: "Footwear",
	},
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(products)
}

func main() {
	http.HandleFunc("/api/products", getProducts)

	println("Server running on http://localhost:8000")

	http.ListenAndServe(":8000", nil)
}
