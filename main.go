package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	db2 "go-hexagonal/adapters/db"
	"go-hexagonal/application"
)

func main() {
	//conexão db
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Exemplo", 30)
	
	productService.Enable(product)
}
