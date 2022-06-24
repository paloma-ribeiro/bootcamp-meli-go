## Introdução

- [Site da linguagem Go](https://golang.com)
- [Site com a documentação de Go](https://go.dev/doc/)

### Características da linguagem Go

- Open Source com objetivo de tornar os programadores mais produtivos
- Expressiva, concisa, limpa e eficiente
- Compilação rápida e trabalha com garbage collection
- Estaticamente tipada e compilada
- Compilada em apenas um arquivo binário
- Framework de testes e profiling nativos
- Detecção de Race conditions
- Deploy simples
- Baixa curva de aprendizado

- Foi criada para suprir necessidades do Google
    - Projeto inicial em setembro de 2007
    - Versão 1.0, lançada em 2012
    - A partir da versão 1.5, seu compilador foi reescrito em Go

- Criadores:
    - Robert Griesemar - V8
    - Rob Pike - Unix & Utf-8
    - Ken Thompson - Unix & Utf-8

### Introdução a linguagem Go

- Cada arquivo em Go pertence a um pacote, sendo necessário declara-los:
  
    > package [nome_do_pacote]

#### Regras de nomenclatura de variável

- Deve começar com uma letra
- Não pode começar com numero
- Não podem conter espaços
- Se o nome da variável começar com letra minuscula:
    - Só poderá ser acessada dentro do pacote atual e serão consideradas como variáveis não exportadas
- Se o nome da variável começar com letra maiuscula:
    - Poderá ser acessada por outros pacotes e serão consideradas variáveis exportadas
- Padrão camelcase. Ex: empName, EmpName, etc

### Observações

- Para instalação de pacotes:
  
  - Ex: go get github.com/google/uuid
  
    > go get [caminho_do_pacote]
    
- Para utilização de pacotes instalados, basta realizar o import dentro do código
  
  - Ex:
  
    > import "github.com/google/uuid"

- Para executar um arquivo go pelo terminal:
  
    > go run [nome_do_arquivo]

- Para iniciar uma aplicação. Executar no terminal:
  
    > go mod init github.com/paloma-ribeiro/bootcamp-meli-go

### Instalação do Testify

> go get github.com/stretchr/testify

### Instalação do Swagger em Go

> go get -u github.com/swaggo/swag/cmd/swag
> go get -u github.com/swaggo/files
> go get -u github.com/swaggo/gin-swagger

#### Gerar documentação com swagger

> swag init -g cmd/server/main.go
