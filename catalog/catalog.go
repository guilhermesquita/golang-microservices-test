package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"

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

var productUrl string

func init() {
	productUrl = os.Getenv("PRODUCT_URL")
}

func loadProducts() []Product {
	response, err := http.Get(productUrl + "/products")
	if err != nil {
		fmt.Println("Error fetching products:", err.Error())
	}

	data, _ := io.ReadAll(response.Body)
	fmt.Println(string(data))
	var products Products
	json.Unmarshal(data, &products)
	return products.Products
}

func listProducts(w http.ResponseWriter, r *http.Request) {
	products := loadProducts()
	template := template.Must(template.ParseFiles("templates/catalog.html"))
	template.Execute(w, products)
}

func showProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response, err := http.Get(productUrl + "/products/" + vars["id"])
	if err != nil {
		fmt.Println("Error fetching product:", err.Error())
	}

	data, _ := io.ReadAll(response.Body)

	var product Product
	json.Unmarshal(data, &product)

	template := template.Must(template.ParseFiles("templates/view.html"))
	template.Execute(w, product)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", listProducts)
	router.HandleFunc("/product/{id}", showProduct)
	http.ListenAndServe(":8080", router)
}
