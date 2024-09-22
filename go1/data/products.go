package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

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

type Products []Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func (p *Products) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func GetProducts() Products {
	cnt := 0
	for _, p := range testProductList {
		if p.DeletedOn == "" {
			cnt++
		}
	}

	result := make(Products, cnt)
	i := 0
	for _, p := range testProductList {
		if p.DeletedOn == "" {
			result[i] = p
			i++
		}
	}
	return result
}

func AddProduct(p *Product) int {
	newId := len(testProductList) + 1
	p.ID = newId
	p.CreatedOn = time.Now().UTC().String()
	p.UpdatedOn = time.Now().UTC().String()
	p.DeletedOn = ""
	testProductList = append(testProductList, *p)
	return newId
}

var ErrorNotFoundProduct = fmt.Errorf("product not found")

func GetProductById(id int) (Product, error) {
	for _, product := range testProductList {
		if product.ID == id && product.DeletedOn == "" {
			return product, nil
		}
	}
	return Product{}, ErrorNotFoundProduct
}

func UpdateProduct(p Product) error {
	for index, product := range testProductList {
		if product.ID == p.ID && product.DeletedOn == "" {
			p.CreatedOn = product.CreatedOn
			p.UpdatedOn = time.Now().UTC().String()
			p.DeletedOn = ""
			testProductList[index] = p
			return nil
		}
	}
	return ErrorNotFoundProduct
}

func UpdateProductById(id int, p Product) error {
	p.ID = id
	return UpdateProduct(p)
}

func DeleteProductById(id int) error {
	for index, product := range testProductList {
		if product.ID == id && product.DeletedOn == "" {
			testProductList[index].DeletedOn = time.Now().UTC().String()
			return nil
		}
	}
	return ErrorNotFoundProduct
}

var testProductList = []Product{
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
