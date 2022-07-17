package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go_inventory_api/domain/entities"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	product := entities.Product{
		Title: "Example Product",
	}
	product.Base.Id = "1234"

	resp, err := json.Marshal(product)
	if err != nil {
		log.Default().Println(err)
	}
	w.Write(resp)
}
