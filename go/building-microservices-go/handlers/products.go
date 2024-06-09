package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prateek041/microservices/data"
)

type Product struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	products := data.GetProducts()
	err := products.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal `json:", http.StatusInternalServerError)
	}
}

func (p *Product) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	// create an empty product struct
	prod := &data.Product{}
	// Unmarshal the received data into the prod struct.
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to Unmarshal json", http.StatusInternalServerError)
	}
	data.AddProducts(prod)
}

func (p *Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Put Products")

	prm := mux.Vars(r)
	id, err := strconv.Atoi(prm["id"])
	if err != nil {
		http.Error(rw, "Unable to convert Id", http.StatusBadRequest)
	}

	prod := &data.Product{}
	err = prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not Foound", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
