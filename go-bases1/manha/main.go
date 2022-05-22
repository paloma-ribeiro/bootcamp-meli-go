package main

import "fmt"

/*
Exercício 1 - Imprimindo o nome na tela

1. Crie uma aplicação que tenha uma variável “nome” e outra “idade”.
2. Imprima no terminal o valor de cada variável.
*/

func imprimeNomeIdade(nome string, idade int) {
	println("Nome:", nome, "\nIdade:", idade)
}

/*
Exercício 2 - Clima

Uma empresa de meteorologia quer ter um sistema onde possa ter a temperatura, umidade e
pressão atmosférica de diferentes lugares.
1. Declare 3 variáveis especificando o tipo de dado delas, como valor elas devem
possuir a temperatura, umidade e pressão atmosférica de onde você se encontra.
2. Imprima os valores no console.
3. Quais tipos de dados serão atribuídos a essas variáveis?
*/

func imprimeClima(temperatura float32, umidade float32, pressaoAtm float32) {
	fmt.Println("Temperatura em Celsius:", temperatura, "\nUmidade em %:", umidade, "\nPressão Atmosférica em atm:", pressaoAtm)
}

/*
Exercício 3 - Declaração de variáveis

Um professor de programação está corrigindo as avaliações de seus alunos na disciplina de
Programação I para fornecer-lhes o feedback correspondente. Um dos pontos do exame é
declarar diferentes variáveis.
Ajude o professor com as seguintes questões:
1. Verifique quais dessas variáveis declaradas pelo aluno estão corretas.
2. Corrigir as incorrectas.

var 1nome string
var sobrenome string
var int idade
1sobrenome := 6
var licenca_para_dirigir = true
var estatura da pessoa int
quantidadeDeFilhos := 2
*/

func declaracaoVariaveis() {
	var nome string
	var sobrenome string
	var idade int
	sobrenome1 := 6
	var LicencaParaDirigir bool = true
	var estaturaDaPessoa int
	quantidadeDeFilhos := 2

	fmt.Println(nome, sobrenome, idade, sobrenome1, LicencaParaDirigir, estaturaDaPessoa, quantidadeDeFilhos)

}

/*
Exercício 4 - Tipos de dados

Um estudante de programação tentou fazer declarações de variáveis com seus respectivos
tipos em Go mas teve várias dúvidas ao fazê-lo. A partir disso, ele nos deu seu código e
pediu a ajuda de um desenvolvedor experiente que pode:
Revisar o código e realizar as devidas correções.

var sobrenome string = "Silva"
var idade int = "25"
boolean := "false";
var salario string = 4585.90
var nome string = "Fellipe"
*/

func tipoDados() {
	var sobrenome string = "Silva"
	var idade int = 25
	boolean := false
	var salario float32 = 4585.90
	var nome string = "Fellipe"

	fmt.Println(sobrenome, idade, boolean, salario, nome)
}

func main() {
	// Exercício 1 - Imprimindo o nome na tela
	imprimeNomeIdade("Paloma Ribeiro", 23)

	// Exercício 2 - Clima
	imprimeClima(23, 12, 148)

	// Exercício 3 - Declaração de variáveis
	declaracaoVariaveis()

	// Exercício 4 - Tipos de dados
	tipoDados()
}
