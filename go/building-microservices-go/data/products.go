package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product defines what a single product looks like
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJson(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

// Products is a collection of Product
type Products []*Product

// Get Products returns a list of Products.
func GetProducts() Products {
	return ProductList
}

func AddProducts(product *Product) {
	product.ID = getNextId()
	ProductList = append(ProductList, product)
}

func getNextId() int {
	product := ProductList[len(ProductList)-1]
	return product.ID + 1
}

var ErrProductNotFound = fmt.Errorf("Product Not Found")

func UpdateProduct(id int, prod *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	prod.ID = id
	ProductList[pos] = prod

	return nil
}

func findProduct(id int) (*Product, int, error) {
	for index, prod := range ProductList {
		if prod.ID == id {
			return prod, index, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// Sample data
var ProductList = Products{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
