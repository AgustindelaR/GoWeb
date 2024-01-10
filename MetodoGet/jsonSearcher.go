package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

func jsonReadProducts(slice *[]Product) {
	file, err := os.Open("products.json")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(slice) // guardo todos los productos
	if err != nil {
		fmt.Println("Error al decodificar el archivo JSON:", err)
		return
	}
}
