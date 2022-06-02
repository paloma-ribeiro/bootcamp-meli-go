# Composição

O conceito de herança não existe em Go, mas temos uma composição que usa uma estrutura semelhante, conhecido como embedding structs.

Exemplo:

```
type Veiculo struct {
    km   float64
    hora float64
}

func (v Veiculo) detalhe() {
    fmt.Printf("km:\t%f\nhora:\t%fn", v.km, v.hora)
}

type Automovel struct {
    v Veiculo
}

func (a *Auto) Correr(minutos int) {
    a.v.hora = float64(minutos) / 60
    a.v.km = a.v.hora * 100
}

func (a *Auto) Detalhe() {
    fmt.Println("\nV:tAutomovel")
    a.v.detalhe()
}

type Moto struct {
    v Veiculo
}

func (m *Moto) Correr(minutos int) {
    m.v.hora = float64(minutos) / 60
    m.v.km = a.v.hora * 80
}

func (m *Moto) Detalhe() {
    fmt.Println("\nV:tMoto")
    m.v.detalhe()
}

func main() {
    automovel := Automovel{}
    auto.Correr(360)
    auto.Detalhe()

    moto := Moto{}
    moto.Correr(360)
    moto.Detalhe()
}
```
