package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ProductCode, Name, Description, ImgURL string
	Weight, Price float64
	Stock int
}

var db *sql.DB
var err error

func getProduct(productCode string) (err error) {
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/go_db")
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	var p Product

	err = db.QueryRow("SELECT productCode, name, description, imageURL, stock, weight, price " + 
		"FROM products " + 
		"WHERE productCode = ? ", productCode).Scan(&p.ProductCode, &p.Name, &p.Description, &p.ImgURL, &p.Stock, &p.Weight, &p.Price)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Product Code: %s\n Image URL: %s\n Product Name: %s\n Description: %s\n Price: $%.2f\n Weight: %.2f\n " +
		"Number in stock: %d\n", p.ProductCode, p.ImgURL, p.Name, p.Description, p.Price, p.Weight, p.Stock)

	return err
}

func main() {
	getProduct("g43")
} 