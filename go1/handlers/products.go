package handlers

import (
	"go1/data"
	"log"
	"net/http"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (handler *Products) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		handler.getTestProductsList(writer, request)
		return
	}

	if request.Method == http.MethodPost {
		handler.addProduct(writer, request)
		return
	}

	http.Error(
		writer,
		"Only GET method is allowed",
		http.StatusMethodNotAllowed,
	)
}

func (handler *Products) getTestProductsList(writer http.ResponseWriter, request *http.Request) {
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

	// TODO: adding product to some storage logic
	handler.logger.Printf("Adding Product %#v", product)
	// ...
}
