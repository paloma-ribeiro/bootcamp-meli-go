package main

import "fmt"

/*
Exercício 1 - Impostos de salário

Uma empresa de chocolates precisa calcular o imposto de seus funcionários no momento de
depositar o salário, para cumprir seu objetivo será necessário criar uma função que retorne o
imposto de um salário.
Temos a informação que um dos funcionários ganha um salário de R$50.000 e será
descontado 17% do salário. Um outro funcionário ganha um salário de $150.000, e nesse
caso será descontado mais 10%.
*/

func calcularImpostoSalario(salario float64) float64 {
	var imposto float64 = 0
	if salario <= 50000 {
		imposto = salario * 0.17
	} else {
		imposto = salario * 0.27
	}
	return imposto
}

/*
Exercício 2 - Calcular média
Um colégio precisa calcular a média das notas (por aluno). Precisamos criar uma função na
qual possamos passar N quantidade de números inteiros e devolva a média.
Obs: Caso um dos números digitados seja negativo, a aplicação deve retornar um erro
*/

func calcularMedia(notas ...int) int {
	var soma int = 0
	media := 0

	for _, nota := range notas {
		soma += nota
	}

	media = soma / len(notas)

	return media
}

func main() {
	// Exercicio 1 - Impostos de salário
	fmt.Println("Imposto que será cobrado do salário do funcionário:", calcularImpostoSalario(60000))
	fmt.Println("Imposto que será cobrado do salário do funcionário:", calcularImpostoSalario(40000))
	fmt.Println("Imposto que será cobrado do salário do funcionário:", calcularImpostoSalario(1600000))

	// Exercício 2 - Calcular Média
	fmt.Println("Media das notas do aluno:", calcularMedia(1, 2, 3))
	fmt.Println("Media das notas do aluno:", calcularMedia(6, 5, 4, 3, 2))
}
