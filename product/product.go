package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Product struct {
	Uuid    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

type Products struct {
	Products []Product
}

func loadData() []byte {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err.Error())
		return nil
	}
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	return data
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadData()
	w.Write(products)
}

func ListProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := loadData()

	var products Products
	json.Unmarshal(data, &products)

	for _, value := range products.Products {
		if value.Uuid == vars["id"] {
			product, _ := json.Marshal(value)
			w.Write(product)
		}
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/products", ListProducts)
	router.HandleFunc("/products/{id}", ListProductById)
	http.ListenAndServe(":8081", router)

	fmt.Println(string(loadData()))
}
