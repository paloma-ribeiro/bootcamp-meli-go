# Estruturas (Structs)

## O que é uma estrutura?

Uma coleção de campos de dados

```
type Pessoa struct {
    Nome       string
    Genero     string
    Idade      int
    Profissao  string
    Peso       float64
}

Exemplo de como instanciar uma estrutura:
```

```
Exemplo 1:

p1 := Pessoa{"Paloma", 20, "Engenheira", 70.2}

Exemplo 2:

p2 := Pessoa {
    Nome: "Vitor",
    Idade: 30,
    Profissao: "Arquiteto",
    Peso: 77,
}

Exemplo 3 - Definindo uma estrutura vazia:

var p3 Pessoa
```

Para acessar um campo da estrutura:

```
p2.Peso
```

Para atribuir ou modificar um valor a um campo da estrutura:

```
p2.Peso = 70
```

Declarar estruturas como campo dentro de outra estrutura:

```
type Preferencias struct {
    Comidas string
    Filme string
    Series string
    Animes string
    Esportes string
}

type Pessoa struct {
    Nome string
    Idade int
    Profissao string
    Peso float64
    Gostos Preferencias
}

p1 := Pessoa{"Paloma", 37, "Engenheira", 65.5, Preferencias{"frango", "titanic", "The Big Bang Theory", "Death Note", "Natação"}}
```

Acessando a estrutura 'Gostos' dentro de 'Pessoa':

```
p2.Gostos.Esportes
```

## Rótulos de estruturas

Dentro das estruturas podemos definir rótulos que se referem a cada um dos campos:

```
type MinhaEstrutura struct {
    Campo1 string `meuRotulo:"valor"`
    Campo2 string `meuRotulo:"valor"`
    Campo3 string `meuRotulo:"valor"`
}
```

Através de tags, podemos especificar o nome de cada campo no formato JSON.

```
import (
    "encoding/json"
)

type Pessoa struct {
    Nome      string `json:"nome"`
    Sobrenome string `json:"sobrenome"`
    Telefone  string `json:"telefone"`
}

p := Pessoa{"Ana", "Banana", "39234567"}
meuJSON, err := json.Marshal(p)

fmt.Println(string(meuJSON))
fmt.Println(err)
```

Definindo uma estrutura com rótulos personalizados:

```
import (
    "reflect"
)

type Pessoa struct {
    Nome      string `bd:"nome"`
    Sobrenome string `bd:"sobrenome"`
    Telefone  string `bd:"telefone"`
}

pessoa := Pessoa{}
p := reflect.TypeOf(pessoa)
```

Verificar o nome e tipo da estrutura:

```
fmt.Println("Type: ", p.Name())
fmt.Println("Kind: ", p.Kind())
```

Método NumField para obter o número de campos de uma estrutura:

```
for i := 0; i < p.NumField(); i++ {

}
```

Método Field para obter o campo da estrutura passando o índice como parâmetro:

```
for i := 0; i < p.NumField(); i++ {
    field := p.Field(i)
}
```

Para acessar o valor do rótulo definido:

```
for i := 0; i < p.NumField(); i++ {
    field := p.Field(i)
    tag := field.Tag.Get("bd")
}
```