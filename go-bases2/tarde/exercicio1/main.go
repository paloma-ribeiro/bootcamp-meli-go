package main

import "fmt"

// Exercício 1 - Registro de estudantes
type Aluno struct {
	nome         string
	sobrenome    string
	rg           string
	dataAdmissao string
}

func imprimeAluno(aluno Aluno) {
	fmt.Println("Nome:", aluno.nome)
	fmt.Println("Sobrenome:", aluno.sobrenome)
	fmt.Println("RG:", aluno.rg)
	fmt.Println("Data de Admissão:", aluno.dataAdmissao)
}

func main() {
	var aluno = Aluno{"Paloma", "Ribeiro", "XX.XXX.XXX-X", "22-05-2022"}

	imprimeAluno(aluno)
}
