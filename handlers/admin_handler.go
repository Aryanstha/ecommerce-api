package handlers

import (
	"ecommerce-api/models"
	"encoding/json"
	"net/http"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	//save product to database
	models.DB.Exec("INSERT INTO products (name, price, quantity) VALUES (?, ?, ?)",
		product.Name, product.Price, product.Quantity)
	w.WriteHeader(http.StatusCreated)
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	//update product in database
	models.DB.Exec("UPDATE products SET name=?, price=?, quantity=? WHERE id=?",
		product.Name, product.Price, product.Quantity, product.ID)
	w.WriteHeader(http.StatusOK)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	//delete product from database
	models.DB.Exec("DELETE FROM products WHERE id=?", product.ID)
	w.WriteHeader(http.StatusOK)
}
