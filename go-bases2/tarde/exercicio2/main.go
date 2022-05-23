package main

import "fmt"

// Exercício 2 - Produtos de e-commerce

const (
	pequeno = "Pequeno"
	medio   = "Médio"
	grande  = "Grande"
)

type loja struct {
	produtos []string
}

type produto struct {
	tipoProduto string
	nome        string
	preco       float64
}

type Produto interface {
	CalcularCusto() float64
}

type Ecommerce interface {
	Total() float64
	Adicionar()
}

func novoProduto(p produto) {
	p = produto{"Pequeno", "Sapato", 34.00}
	fmt.Println(p)
}
