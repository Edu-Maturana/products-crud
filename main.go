package main

import (
	"fmt"
	c "go-crud/controllers"
	env "go-crud/database"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := env.GoDotEnvVar("PORT")

	http.HandleFunc("/api/products", c.GetProducts)
	http.HandleFunc("/api/products/", c.CreateProduct)
	http.HandleFunc("/api/products/update/", c.UpdateProduct)
	http.HandleFunc("/api/products/delete/", c.DeleteProduct)

	http.ListenAndServe(":"+port, nil)
	fmt.Println("Server started on port " + port)
}
