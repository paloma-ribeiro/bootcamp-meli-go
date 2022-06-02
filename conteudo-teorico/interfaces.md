# Interfaces

## O que são interfaces?

É uma forma de definir métodos que devem ser usados, mas sem defini-los, e servem para fornecer modularidade à linguagem


Exemplo:

```
type circulo struct {
    raio float64
}

func (c circle) area() float64 {
    return math.Pi * c.raio * c.raio
}

func (c circle) perimetro() float64 {
    return 2 * math.Pi * c.raio
}

type geometria interface {
    area() float64
    perimetro() float64
}

type trianfulo struct {
    largura, altura float64
}

func (t triangulo) area() float64 {
    return t.largura * t.altura
}

func (t rect) perimetro() float64 {
    return 2*t.largura + 2*t.altura
}

func detalhe(g geometria) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perimetro())
}

func main() {
    t := triangulo{largura: 3, altura: 4}
    c := circulo{raio: 5}
    detalhes(t)
    detalhes(c)
}
```
