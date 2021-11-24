package models

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Image       string `json:"image"`
}

func (u *Product) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
