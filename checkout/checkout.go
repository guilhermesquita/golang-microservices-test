package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var productUrl string

func init() {
	productUrl = os.Getenv("PRODUCT_URL")
}

type Product struct {
	Uuid    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

type Order struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	ProductId string `json:"product_id"`
}

func displayCheckout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response, err := http.Get(productUrl + "/products/" + vars["id"])
	if err != nil {
		fmt.Println("Error fetching product:", err.Error())
	}

	data, _ := io.ReadAll(response.Body)

	var product Product
	json.Unmarshal(data, &product)

	template := template.Must(template.ParseFiles("templates/checkout.html"))
	template.Execute(w, product)
}

func finish(w http.ResponseWriter, r *http.Request) {
	var order Order
	order.Name = r.FormValue("name")
	order.Email = r.FormValue("email")
	order.Phone = r.FormValue("phone")
	order.ProductId = r.FormValue("product_id")

	data, _ := json.Marshal(order)

	fmt.Sprintln(string(data))
	w.Write([]byte("Processou"))

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/finish", finish)
	router.HandleFunc("/{id}", displayCheckout)
	http.ListenAndServe(":8082", router)
}
