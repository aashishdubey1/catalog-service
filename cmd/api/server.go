package main

import (
	"fmt"
	"net/http"

	"github.com/aashishdubey1/catalog-service/internal/db"
	"github.com/aashishdubey1/catalog-service/internal/product"
	"github.com/go-chi/chi/v5"
)

func main() {
	conn, err := db.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	// do once
	// db.CreateTables(conn)
	// db.SeedProducts(conn)

	r := chi.NewRouter()

	productRepo := product.NewPgProductRepo(conn)
	productService := product.NewProductService(productRepo)
	productHandler := product.NewProductHandler(productService)

	r.Mount("/products", productHandler.Routes())

	http.ListenAndServe(":8080", r)
}
