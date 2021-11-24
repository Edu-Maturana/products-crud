package controllers

import (
	"encoding/json"
	"fmt"
	db "go-crud/database"
	m "go-crud/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/nu7hatch/gouuid"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	conn := db.DBConnection()

	rows, err := conn.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}

	var products []m.Product
	for rows.Next() {
		var product m.Product
		err = rows.Scan(&product.Id, &product.Name, &product.Category, &product.Description, &product.Price, &product.Image)
		if err != nil {
			panic(err.Error())
		}
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
	var product m.Product // in models/products.go
	err := decoder.Decode(&product)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	response, err := product.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Connect to the database
	conn := db.DBConnection()

	// Asign the id
	id, err := uuid.NewV4()
	if err != nil {
		panic(err.Error())
	}
	product.Id = id.String()

	// Prepare the statement
	insert, err := conn.Prepare("INSERT INTO products(id, name, category, description, price, image) VALUES('" + product.Id + "','" + product.Name + "','" + product.Category + "','" + product.Description + "','" + product.Price + "', '" + product.Image + "')")
	if err != nil {
		panic(err.Error())
	}

	// Execute the statement
	insert.Exec()

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	decoder := json.NewDecoder(r.Body)
	var product m.Product // in models/products.go
	err := decoder.Decode(&product)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	response, err := product.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Connect to the database
	conn := db.DBConnection()

	// Prepare the statement
	update, err := conn.Prepare("UPDATE products SET name = ?, category = ?, description = ?, price = ?, image = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	// Execute the statement
	update.Exec(product.Name, product.Category, product.Description, product.Price, product.Image, productId)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	fmt.Println(productId)

	// Connect to the database
	conn := db.DBConnection()
	delete, err := conn.Prepare("DELETE FROM products WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}

	// Execute the statement
	delete.Exec(productId)
	w.Write([]byte("Product deleted"))
}
