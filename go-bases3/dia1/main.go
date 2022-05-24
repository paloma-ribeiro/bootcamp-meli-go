package main

import (
	"fmt"
	"os"
)

// Exercício 1 - Guardar arquivo

func guardarArquivo() {
	produtosComprados := []byte("Banana; Maça; Laranja; Goiaba")
	err := os.WriteFile("myFile.csv", produtosComprados, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(produtosComprados))
	}
}

func main() {
	guardarArquivo()
}
