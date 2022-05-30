# Enunciado dos Exercícios

## Exercício 1 - Estruturar um JSON

Dependendo do tema escolhido, gere um JSON que atenda as seguintes chaves de acordo
com o tema.

Os produtos variam por id, nome, cor, preço, estoque, código (alfanumérico), publicação
(sim-não), data de criação.

Os usuários variam por id, nome, sobrenome, e-mail, idade, altura, ativo (sim-não), data de
criação.

Transações: id, código da transação (alfanumérico), moeda, valor, emissor (string), receptor
(string), data da transação.

1. Dentro da pasta go-web crie um arquivo theme.json, o nome tem que ser o tema
escolhido, ex: products.json.

2. Dentro dele escreva um JSON que permite ter uma matriz de produtos, usuários ou
transações com todas as suas variantes.

## Exercício 2 - Olá {nome}

1. Crie dentro da pasta go-web um arquivo chamado main.go

2. Crie um servidor web com Gin que retorne um JSON que tenha uma chave
“mensagem” e diga Olá seguido do seu nome.

3. Acesse o end-point para verificar se a resposta está correta.

## Exercício 3 - Listar Entidade

Já tendo criado e testado nossa API que nos recebe, geramos uma rota que retorna uma lista
do tema escolhido.

1. Dentro do "main.go", crie uma estrutura de acordo com o tema com os campos
correspondentes.

2. Crie um endpoint cujo caminho é /thematic (plural). Exemplo: “/products”

3. Crie uma handler para o endpoint chamada "GetAll".

4. Crie um slice da estrutura e retorne-a por meio de nosso endpoint.