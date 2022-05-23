package main

import "fmt"

/*
Exercício 1 - Letras de uma palavra

A Real Academia Brasileira quer saber quantas letras tem uma palavra e então ter cada uma
das letras separadamente para soletrá-la. Para isso terão que:
1. Crie uma aplicação que tenha uma variável com a palavra e imprima o número de
letras que ela contém.
2. Em seguida, imprimi cada uma das letras.
*/

func imprimeLetra(palavra string) {

	fmt.Println("Quantidade de Letras da palavra:", len(palavra))

	for _, letra := range palavra {
		fmt.Println(string(letra))
	}
}

/*
Exercício 2 - Empréstimo

Um banco deseja conceder empréstimos a seus clientes, mas nem todos podem acessá-los.
Para isso, possui certas regras para saber a qual cliente pode ser concedido. Apenas
concede empréstimos a clientes com mais de 22 anos, empregados e com mais de um ano
de atividade. Dentro dos empréstimos que concede, não cobra juros para quem tem um
salário superior a US$ 100.000.
É necessário fazer uma aplicação que possua essas variáveis e que imprima uma mensagem
de acordo com cada caso.
Dica: seu código deve ser capaz de imprimir pelo menos 3 mensagens diferentes.
*/

type Cliente struct {
	nome          string
	idade         int
	empregado     bool
	tempoTrabalho float32
	salario       float32
}

func verificaCliente(cliente Cliente) []bool {

	var estadoCliente = []bool{}

	if cliente.idade > 22 {
		estadoCliente = append(estadoCliente, true)
	} else {
		estadoCliente = append(estadoCliente, false)
	}
	if cliente.empregado == true {
		estadoCliente = append(estadoCliente, true)
	} else {
		estadoCliente = append(estadoCliente, false)
	}
	if cliente.tempoTrabalho > 1 {
		estadoCliente = append(estadoCliente, true)
	} else {
		estadoCliente = append(estadoCliente, false)
	}
	if cliente.salario > 100000 {
		estadoCliente = append(estadoCliente, true)
	} else {
		estadoCliente = append(estadoCliente, false)
	}
	return estadoCliente
}

func verificaEmprestimo(estadoCliente []bool) {

	if estadoCliente[0] && estadoCliente[1] && estadoCliente[2] == true {
		if estadoCliente[3] == true {
			fmt.Println("Cliente apto ao emprestimo sem cobrança de juros")
		} else {
			fmt.Println("Cliente apto ao emprestimo com cobrança de juros")
		}
	} else {
		fmt.Println("Cliente inapto ao emprestimo")
	}

}

/*
Exercício 3 - A que mês corresponde?

Faça uma aplicação que contenha uma variável com o número do mês.
1. Dependendo do número, imprima o mês correspondente em texto.
2. Ocorre a você que isso pode ser resolvido de maneiras diferentes? Qual você
escolheria e porquê?
*/

func imprimeMes(numeroMes int) {
	var expressao bool = true

	switch expressao {
	case numeroMes == 1:
		println("Janeiro")
	case numeroMes == 2:
		println("Fevereiro")
	case numeroMes == 3:
		println("Março")
	case numeroMes == 4:
		println("Abril")
	case numeroMes == 5:
		println("Maio")
	case numeroMes == 6:
		println("Junho")
	case numeroMes == 7:
		println("Julho")
	case numeroMes == 8:
		println("Agosto")
	case numeroMes == 9:
		println("Setembro")
	case numeroMes == 10:
		println("Outubro")
	case numeroMes == 11:
		println("Novembro")
	case numeroMes == 12:
		println("Dezembro")
	default:
		println("Número inválido!")
	}
}

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

var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

func imprimeFuncionario(nome string) {

	for chave, valor := range employees {
		if chave == nome {
			fmt.Println("Nome:", chave, "Idade:", valor)
		}
	}
}

func verificaIdadeFuncionarios(idade int) {
	count := 0

	for _, valor := range employees {
		if valor > idade {
			count++
		}
	}
	fmt.Println(count, "funcionários possuem mais de", idade, "anos")
}

func addFuncionario(nome string, idade int) {
	employees[nome] = idade
	fmt.Println(employees)
}

func excluiFuncionario(nome string) {
	delete(employees, nome)
	fmt.Println(employees)
}

func main() {
	// Exercício 1 - Letras de uma palavra
	imprimeLetra("Bootcamp Meli")

	// Exercício 2 - Empréstimo
	cliente1 := Cliente{"Paloma Ribeiro", 25, true, 2, 5000000}
	cliente2 := Cliente{"Joana da Silva", 18, true, 0.5, 50000}
	cliente3 := Cliente{"Romulo Andrade", 23, true, 3, 50000}

	resultadoCliente1 := verificaCliente(cliente1)
	resultadoCliente2 := verificaCliente(cliente2)
	resultadoCliente3 := verificaCliente(cliente3)

	verificaEmprestimo(resultadoCliente1)
	verificaEmprestimo(resultadoCliente2)
	verificaEmprestimo(resultadoCliente3)

	// Exercício 3 - A que mês corresponde?
	imprimeMes(7)
	imprimeMes(12)
	imprimeMes(14)

	// Exercício 4 - Quantos anos tem...
	imprimeFuncionario("Benjamin")
	verificaIdadeFuncionarios(21)
	addFuncionario("Julio", 26)
	excluiFuncionario("Pedro")
}
