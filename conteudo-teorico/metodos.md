# Métodos

## O que são métodos?

Go não possui classes. No entanto, os métodos podem ser definidos em tipos de dados.

Um método é uma função com um argumento de receptor especial. O receptor aparece em sua própria lista de argumentos entre a palavra-chave func e o nome do método

```
type Circulo struct {
    raio float64
}

func (v MinhaEstrutura) metodo(){}

Declarando o método área:

func (c Circulo) area() float64 {
    return math.Pi * c.raio * c.raio
}

Declarando o método perímetro:

func (c Circulo) perimetro() float64 {
    return 2 * math.Pi * c.raio
}
```

Para modificar variáveis da estrutura no método, deve-se passar o valor como referência e indica-lo como ponteiro. Caso contrário, a variável não será modificada ao sair do escopo do método

```
func (c *Circulo) setRaio(r float64) {
    c.raio = r
}
```

Para executar os métodos:

```
func main() {
    c := Circulo{raio: 5}

    fmt.Println(c.area())
    fmt.Println(c.perimetro())
    
    c.setRaio(10)
    fmt.Println(c.area())
    fmt.Println(c.perimetro())

}
```