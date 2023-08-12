package main

import (
	"ecommerce-api/handlers"
	"ecommerce-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDB()
	defer models.DB.Close()

	r := mux.NewRouter()

	// Admin routes
	r.HandleFunc("/admin/product", handlers.AddProduct).Methods("POST")
	r.HandleFunc("/admin/product/{id}", handlers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/admin/product/{id}", handlers.DeleteProduct).Methods("DELETE")

	// Client routes
	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
