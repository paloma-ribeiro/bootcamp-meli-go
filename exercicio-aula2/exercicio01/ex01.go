package exercicio01

import "fmt"

/*
Exercício 1 - Imprimindo o nome na tela
1. Crie uma aplicação que tenha uma variável “nome” e outra “idade”.
2. Imprima no terminal o valor de cada variável.
*/

func imprimeNomeIdade() {
	nome := "Paloma Ribeiro"
	idade := 20

	fmt.Printf("Nome: %s, Idade: %s", nome, idade)
}
