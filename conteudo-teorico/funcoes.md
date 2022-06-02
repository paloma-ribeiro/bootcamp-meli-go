# Funções 

## O que é uma função?

Pedaço de código que executa uma tarefa específica

## Estrutura de uma função

A funçao recebe um, muitos parâmetros ou nenhum, e pode ou não retornar um valor

```
func minhaFuncao(parâmetros) {

}
```

Exemplo:

```
func verificarVariavel(numero int) {
    if numero < 0 {
        fmt.Println("Número negativo")
    } else {
        fmt.Println("Número positivo")
    }
}
```
Função com mais de um parâmetro

```
func minhaFuncao(valor1 float64, valor2 float64)

func minhaFuncao(valor1, valor2 float64)
```

Retorno de valores em funções

```
func soma(valor1, valor2 float64) float64 {
    return valor1 + valor2
}
```

## Ellipsis

Notação de reticências utilizado para indicar que uma função receberá um número dinâmico de parâmetros

```
func minhaFuncao(valores ...float64) float64 {

}
```
Obs: ao chamar essa função, podemos passar a quantidade de valores que queremos, sempre do mesmo tipo de dado, e a função receberá os parâmetros como se fossem um array.

Exemplo:

```
func soma(valores ...float64) float64 {
    var resultado float64

    for _, valor := range valores {
        resultado += valor
    }
    return resultado
}

soma(2, 3, 5, 4, 15)
```

Também é possivel passar outros parâmetros, de tipos diferentes, mas o parâmetro de reticências de ficar sempre no final

```
func minhaFuncao(valor1 string, valor2 string, valores ...float64) {

}
```

## Multi retorno

Podemos criar funções que retornam mais de um valor

```
func minhaFuncao(valor1, valor2 float54) (float64, string, int, bool) {

}
```

## Retorno de valores nomeados

Também podemos retornar valores nomeados

```
func operaciones(valor1, valor2 float64) (soma float64, subtra float64, multip float64, divis float64)
```

## Retorno de função

Podemos implementar uma função que retorna outra função

```
func minhaFuncao(valor string) func(valor1, valor2 float64) float64
```