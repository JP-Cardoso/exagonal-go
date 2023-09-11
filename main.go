package main

import (
	"database/sql"
	"go-hexagonal/application"
	db2 "go-hexagonal/adapters/db"
	_ "github.com/mattn/go-sqlite3"

)

func main() {
	println("opa")
	//conexão db
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	productService.Create("Product Exemplo", 30)
}
