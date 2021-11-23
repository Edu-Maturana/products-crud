package controllers

import (
	"encoding/json"
	"fmt"
	db "go-crud/database"
	m "go-crud/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	conn := db.DBConnection()

	rows, err := conn.Query("SELECT * FROM productos")
	if err != nil {
		panic(err.Error())
	}

	var products []m.Product
	for rows.Next() {
		var product m.Product
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Image)
		if err != nil {
			panic(err.Error())
		}
		product.ToJson()
		products = append(products, product)
	}

	response, err := json.Marshal(products)
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var user m.Product // in models/products.go
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Connect to the database
	conn := db.DBConnection()

	// Prepare the statement
	insert, err := conn.Prepare("INSERT INTO productos(name, description, price, image) VALUES('" + user.Name + "','" + user.Description + "','" + user.Price + "', '" + user.Image + "')")
	if err != nil {
		panic(err.Error())
	}

	// Execute the statement
	insert.Exec()

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
