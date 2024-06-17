package handlers

import (
	"context"
	"fmt"
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

	prod := r.Context().Value(ProductKey{}).(data.Product)
	data.AddProducts(&prod)
}

func (p *Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Put Products")

	prm := mux.Vars(r)
	id, err := strconv.Atoi(prm["id"])
	if err != nil {
		http.Error(rw, "Unable to convert Id", http.StatusBadRequest)
	}

	prod := r.Context().Value(ProductKey{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not Foound", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type ProductKey struct{}

func (p *Product) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		p.l.Println("[MESSAGE]: Middleware Triggered")
		prod := &data.Product{}
		err := prod.FromJson(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		// validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), ProductKey{}, *prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
