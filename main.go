package main

import (
	"fmt"
	c "go-crud/controllers"
	"net/http"
)

func main() {

	http.HandleFunc("/api/products", c.GetProducts)
	http.HandleFunc("/api/product", c.CreateProduct)
	http.HandleFunc("/api/product/update", c.UpdateProduct)
	http.HandleFunc("/api/products/delete", c.DeleteProduct)

	http.ListenAndServe(":3000", nil)
	fmt.Println("Server started!")
}
