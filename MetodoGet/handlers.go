package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func bienvenidaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Bienvenido a mi API WEB !")
}
func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "pong")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// productos => global variable
	ec := json.NewEncoder(w)
	if err := ec.Encode(products); err != nil {
		fmt.Println(err)
		return
	}
}

func productByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	//Buscar el producto por ID en la slice
	var foundProduct *Product
	for i := range products {
		if products[i].Id == id {
			foundProduct = &products[i]
			break
		}
	}

	if foundProduct != nil {
		json.NewEncoder(w).Encode(foundProduct)
	} else {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
	}
}

func prodctsSearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	priceGtParam := chi.URLParam(r, "priceGt")
	priceGt, err := strconv.ParseFloat(priceGtParam, 64)
	if err != nil {
		http.Error(w, "Valor de precio no válido", http.StatusBadRequest)
		return
	}

	// Filtrar los productos por precio mayor a priceGt
	var filteredProducts []Product
	for _, product := range products {
		if product.Price > priceGt {
			filteredProducts = append(filteredProducts, product)
		}
	}

	json.NewEncoder(w).Encode(filteredProducts)
}
