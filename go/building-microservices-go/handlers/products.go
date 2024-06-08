package handlers

import (
	"log"
	"net/http"

	"github.com/prateek041/microservices/data"
)

type Product struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	// catch all statement
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	products := data.GetProducts()
	err := products.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal `json:", http.StatusInternalServerError)
	}
}

func (p *Product) addProduct(rw http.ResponseWriter, r *http.Request) {
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
