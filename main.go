package main

import (
	c "go-crud/controllers"
)

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/api/product", c.GetProducts)
	server.Handle("POST", "/api/product", c.CreateProduct)
	server.Handle("PUT", "/api/product", c.UpdateProduct)
	server.Handle("DELETE", "/api/product", c.DeleteProduct)
	server.Listen()
}
