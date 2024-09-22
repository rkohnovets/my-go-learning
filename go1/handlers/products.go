package handlers

import (
	"errors"
	"go1/data"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (handler *Products) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler.logger.Printf(
		"Products handler call, '%s' with path '%s'",
		request.Method,
		request.URL.Path,
	)

	// get all products
	if request.Method == http.MethodGet {
		handler.getTestProductsList(writer, request)
		return
	}

	// add new product
	if request.Method == http.MethodPost {
		handler.addProduct(writer, request)
		return
	}

	// update existing product
	if request.Method == http.MethodPut {
		handler.updateProduct(writer, request)
		return
	}

	if request.Method == http.MethodDelete {
		handler.deleteProduct(writer, request)
		return
	}

	http.Error(
		writer,
		"Method not supported",
		http.StatusMethodNotAllowed,
	)
}

func (handler *Products) getTestProductsList(writer http.ResponseWriter, _ *http.Request) {
	productsList := data.GetProducts()
	err := productsList.ToJSON(writer)
	if err != nil {
		handler.logger.Println(
			"Error in 'Products' handler",
			"(stringify products list to json):",
			err,
		)
		http.Error(
			writer,
			"Error",
			http.StatusInternalServerError,
		)
		return
	}
}

func (handler *Products) addProduct(writer http.ResponseWriter, request *http.Request) {
	product := data.Product{}
	err := product.FromJSON(request.Body)
	if err != nil {
		handler.logger.Println(
			"Error in 'Products' handler",
			"(parsing product from json string):",
			err,
		)
		http.Error(
			writer,
			"Error: unable to parse provided json",
			http.StatusBadRequest,
		)
		return
	}

	data.AddProduct(&product)
	handler.logger.Printf("Added Product: %#v", product)
}

func (handler *Products) updateProduct(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	parts := strings.Split(path, "/")
	// for example "/products/1" => "", "products", "1"
	if len(parts) != 3 {
		http.Error(
			writer, "Bad request - invalid url",
			http.StatusBadRequest,
		)
		return
	}

	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		http.Error(
			writer, "Bad request - unable to parse id",
			http.StatusBadRequest,
		)
		return
	}

	newProduct := data.Product{}
	err = newProduct.FromJSON(request.Body)
	if err != nil {
		http.Error(
			writer, "Bad request - unable to parse json",
			http.StatusBadRequest,
		)
		return
	}

	err = data.UpdateProductById(id, newProduct)
	if err != nil && errors.Is(err, data.ErrorNotFoundProduct) {
		http.Error(
			writer, "Bad request - product not found",
			http.StatusBadRequest,
		)
		return
	}
	if err != nil {
		http.Error(
			writer, "Internal server error",
			http.StatusInternalServerError,
		)
		return
	}
}

func (handler *Products) deleteProduct(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	parts := strings.Split(path, "/")
	// for example "/products/1" => "", "products", "1"
	if len(parts) != 3 {
		http.Error(
			writer, "Bad request - invalid url",
			http.StatusBadRequest,
		)
		return
	}

	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		http.Error(
			writer, "Bad request - unable to parse id",
			http.StatusBadRequest,
		)
		return
	}

	err = data.DeleteProductById(id)
	if errors.Is(err, data.ErrorNotFoundProduct) {
		http.Error(
			writer, "Bad request - product not found",
			http.StatusBadRequest,
		)
		return
	}
	if err != nil {
		http.Error(
			writer, "Internal server error",
			http.StatusInternalServerError,
		)
		return
	}
}
