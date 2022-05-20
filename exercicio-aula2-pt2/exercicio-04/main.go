package main

import "fmt"

/*
Exercício 4 - Quantos anos tem...
Um funcionário de uma empresa deseja saber o nome e a idade de um de seus funcionários.
De acordo com o mapa abaixo, ajude a imprimir a idade de Benjamin.

var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

Por outro lado, você também precisa:
- Saiba quantos de seus funcionários têm mais de 21 anos.
- Adicione um novo funcionário à lista, chamado Federico, que tem 25 anos.
- Excluir Pedro do mapa.
*/

func main() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println("Idade de Benjamin", employees["Benjamin"])

	var contador int = 0

	for _, value := range employees {
		if value > 21 {
			contador++
		}
	}

	fmt.Println("Valor de contador: ", contador)

	employees["Frederico"] = 25

	delete(employees, "Pedro")

	fmt.Println(employees)
}
