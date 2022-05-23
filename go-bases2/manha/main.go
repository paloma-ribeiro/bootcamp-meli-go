package main

import (
	"fmt"
	"sort"
)

// Exercício 1 - Impostos de salário
func calculaImpostoSalario(salario float32) float32 {
	if salario >= 50000 {
		return salario * 0.17
	} else if salario >= 150000 {
		return salario * 0.27
	}
	return 0
}

// Exercício 2 - Calcular média
func calculaMediaNotas(notas ...int) float32 {
	soma := 0
	media := 0

	for _, nota := range notas {
		soma += nota
	}

	media = soma / len(notas)

	return float32(media)
}

// Exercício 3 - Calcular salário
func calculaSalario(horasTrabalhadas float64, categoria string) float64 {
	var salario float64 = 0
	adicional1 := 0.20
	adicional2 := 0.50

	if categoria == "C" {
		salario = 1000 * horasTrabalhadas
	} else if categoria == "B" {
		if horasTrabalhadas < 160 {
			salario = 1500 * horasTrabalhadas
		} else {
			salario = (1500 * horasTrabalhadas) + (salario * adicional1)
		}
	} else if categoria == "A" {
		if horasTrabalhadas < 160 {
			salario = 3000 * horasTrabalhadas
		} else {
			salario = (3000 * horasTrabalhadas) + (salario * adicional2)
		}
	}

	return salario
}

// Exercício 4 - Cálculo de estatísticas
const (
	minimo = "minimo"
	media  = "media"
	maximo = "maximo"
)

func calculaMinimo(valores ...int) {
	valorMinimo := sort.IntSlice(valores)

	println(valorMinimo[0])
}

func calculaMaximo(valores ...int) {
	valorMaximo := sort.IntSlice(valores)

	println(valorMaximo[len(valorMaximo)-1])
}

// Exercício 5 - Cálculo da quantidade de alimento
const (
	cao       = "cao"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarântula"
)

func quantidadeAlimento(animal string, quantidadeAnimal int) float64 {
	if animal == cao {
		return float64(quantidadeAnimal * 10)
	}
	if animal == gato {
		return float64(quantidadeAnimal * 5)
	}
	if animal == hamster {
		return float64(float64(quantidadeAnimal) * 0.25)
	}
	if animal == tarantula {
		return float64(float64(quantidadeAnimal) * 0.15)
	}
	return 0
}

func main() {
	// Exercício 1 - Impostos de salário
	fmt.Println("Imposto:", calculaImpostoSalario(150000))
	fmt.Println("Imposto:", calculaImpostoSalario(50000))
	fmt.Println("Imposto:", calculaImpostoSalario(1500))

	// Exercício 2 - Calcular média
	fmt.Println("Media:", calculaMediaNotas(2, 5, 3, 5, 5))

	// Exercício 3 - Calcular salário
	fmt.Println(calculaSalario(162, "C"))
	fmt.Println(calculaSalario(176, "B"))
	fmt.Println(calculaSalario(172, "A"))

	// Exercício 4 - Cálculo de estatísticas
	calculaMinimo(2, 3, 3, 4, 10, 2, 4, 5)
	calculaMaximo(2, 3, 3, 4, 10, 2, 4, 5)

	// Exercício 5 - Cálculo da quantidade de alimento
	fmt.Println(quantidadeAlimento(cao, 1))
	fmt.Println(quantidadeAlimento(tarantula, 2))
}
