package main

import (
	"fmt"
	"net/http"
)

type produto struct {
	Id          int
	Nome        string
	Cor         string
	Preco       float64
	Estoque     int
	Codigo      int
	Publicacao  bool
	DataCriacao string
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Olá, Mundo!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
