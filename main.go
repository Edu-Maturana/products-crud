package main

import (
	"fmt"
	"go-crud/database"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", CreateProduct)
	fmt.Println("Server running")
	http.ListenAndServe(":8000", nil)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	conn := database.DBConnection()
	insert, err := conn.Prepare("INSERT INTO productos(name, description, price, image) VALUES('Pepperoni', 'Mucho queso, todo muy rico', '15000', 'https://www.pedidosya.cl/blog/wp-content/uploads/sites/2/2018/05/pizza-todas-las-carnes.png')")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec()
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		image := r.FormValue("image")

		conn := database.DBConnection()
		insert, err := conn.Prepare("INSERT INTO productos(name, description, price, image) VALUES('" + name + "','" + description + "','" + price + "','" + image + "')")
		if err != nil {
			panic(err.Error())
		}
		insert.Exec()
	}
}
