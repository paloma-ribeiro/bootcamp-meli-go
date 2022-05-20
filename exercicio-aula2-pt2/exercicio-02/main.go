package main

import (
	"fmt"
	"time"
)

/*
Um banco deseja conceder empréstimos a seus clientes, mas nem todos podem acessá-los.
Para isso, possui certas regras para saber a qual cliente pode ser concedido. Apenas
concede empréstimos a clientes com mais de 22 anos, empregados e com mais de um ano
de atividade. Dentro dos empréstimos que concede, não cobra juros para quem tem um
salário superior a US$ 100.000.
É necessário fazer uma aplicação que possua essas variáveis e que imprima uma mensagem
de acordo com cada caso.
Dica: seu código deve ser capaz de imprimir pelo menos 3 mensagens diferentes.
*/

func main() {
	var (
		nome           = "Joao"
		idade          = 20
		empregado      = true
		dataDeAdmissao = "2022-05-01T00:00:00Z"
		salario        = 100.000
	)

	dataAgora := time.Now()
	dataFormatada, _ := time.Parse(time.RFC3339, dataDeAdmissao)
	tempoTrabalho := dataAgora.Sub(dataFormatada)
	// tempo := tempoTrabalho / 24
	fmt.Println(tempoTrabalho)

	if idade > 22 {
		if empregado == true {
			if salario > 100.000 {
				fmt.Println("O cliente", nome, "está elegivel, tem direito ao emprestimo com taxas de juros menores.")
			} else {
				fmt.Println("O cliente", nome, "está elegivel para fazer emprestimos.")
			}
		} else {
			fmt.Println("O cliente", nome, "não está elegivel, não fornecemos emprestimos para desempregados.")
		}
	} else {
		fmt.Println("O cliente", nome, "não está elegivel, não fornecemos emprestimos para menores de 21 anos.")
	}
}
