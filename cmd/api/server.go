package main

import (
	"fmt"
	"net/http"

	"github.com/aashishdubey1/catalog-service/internal/db"
)

func main() {
	conn, err := db.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	db.CreateTables(conn)
	db.SeedProducts(conn)

	mux := http.NewServeMux()

	http.ListenAndServe(":8080", mux)
}
