package main

import (
	"errors"
	"fmt"
)

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

func calcularMedia(notas ...int) (int, error) {

	var soma int = 0
	media := 0

	for _, nota := range notas {

		if nota < 0 {
			return 0, errors.New("Nota não pode ser negativa")
		} else {
			soma += nota
		}
	}

	media = soma / len(notas)

	return media, nil
}

func main() {
	// Exercicio 1 - Impostos de salário
	fmt.Println("Imposto que será cobrado do salário do funcionário:", calcularImpostoSalario(60000))
	fmt.Println("Imposto que será cobrado do salário do funcionário:", calcularImpostoSalario(40000))
	fmt.Println("Imposto que será cobrado do salário do funcionário:", calcularImpostoSalario(1600000))

	// Exercício 2 - Calcular Média
	resultado, erro := calcularMedia(1, 2, 3)
	resultado1, erro1 := calcularMedia(-1, 2, 3, 8)

	if erro != nil {
		fmt.Println("Erro", erro)
		fmt.Println("Erro", erro1)
	} else {
		fmt.Println("Média do aluno:", resultado)
		fmt.Println("Média do aluno:", resultado1)
	}

}
