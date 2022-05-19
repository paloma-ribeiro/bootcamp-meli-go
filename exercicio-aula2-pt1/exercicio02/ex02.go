package exercicio02

import "fmt"

/*
Exercício 2 - Clima
Uma empresa de meteorologia quer ter um sistema onde possa ter a temperatura, umidade e
pressão atmosférica de diferentes lugares.
1. Declare 3 variáveis especificando o tipo de dado delas, como valor elas devem
possuir a temperatura, umidade e pressão atmosférica de onde você se encontra.
2. Imprima os valores no console.
3. Quais tipos de dados serão atribuídos a essas variáveis?
*/

var temperatura int = 16
var umidade int = 54
var pressaoAtmosferica int = 690

func imprimeDados() {
	fmt.Println("Temperatura: ", temperatura, "Umidade: ", umidade, "Pressão Atmosférica: ", pressaoAtmosferica)
}
