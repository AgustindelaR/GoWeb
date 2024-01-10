package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// url de json https://drive.google.com/file/d/1oZ71o1BCml2EGhAQ31wvtv-RGZzTQjaW/view
var products []Product

func serverInit() {
	router := chi.NewRouter()

	jsonReadProducts(&products) // read and process full json into slice

	router.Get("/", bienvenidaHandler)

	router.Get("/ping", pingHandler)

	router.Get("/products", productsHandler)
	router.Get("/products/{id}", productByIDHandler)
	router.Get("/products/search/{priceGt}", prodctsSearchHandler)

	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

func main() {
	serverInit()
}
