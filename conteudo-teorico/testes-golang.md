# Testes em Go

## O que é teste?

Conjunto de processos, métodos e ferramentas que visam identificar defeitos ou inconsistências no software.

### Teste de caixa preta

Método de teste em que a estrutura interna do projeto ou operação interna é desconhecida pela pessoa que executa o teste.

Seu objetivo é testar a funcionalidade do código, avaliando as respostas e reações (comportamentos) do componente testado em diferentes cenários.

É aplicado para:

- Testes funcionais: validam se o software atende a um determinado requisito
- Testes não funcionais: apenas para medição e avaliação de desempenho, usabilidade e escalabilidade
- Testes de regressão: são executados após qualquer alteração no código, para verificar se esta não afetou o comportamento esperado

### Teste de caixa branca

Método onde quem executa o teste conhece e tem visibilidade sobre o código.

Se concentram nos detalhes processuais do software, de modo que seu design está intimamente ligado ao código-fonte.

É aplicado para testar o fluxo, segurança e estrutura do código, e assim detectar vulnerabilidades, implementação correta de cada método ou função e validar se o fluxo de dados se comporta conforme o esperado, como condicionais, iterações e respostas.

Podem ser aplicados através de:

- Testes unitários
- Testes de integração

## Tipos de testes

### Testes Unitários

Consiste em testar individualmente as funções e métodos das classes, componentes e módulos que são utilizados pelo software.

Principais benefícios: 

- Facilita as alterações de código
- Encontrar bugs
- Fornecem documentação
- Melhoram o design e a qualidade do código

### Testes de integração

Testam a comunicação entre diferentes componentes ou camadas da aplicação.

Seu objetivo é verificar se todos aqueles blocos de código que foram testados de forma unitária, interagem e se comunicam entre si, gerando os resultados esperados.

Eles expoem o funcionamento geral da aplicação, permitindo avaliar seu design, desempenho e comportamento.

### Testes Funcionais

Consiste em validar o comportamento funcional do software, e verificar se ele corresponde exatamente ao resultado esperado a cada cenário.

A funcionalidade deve ser avaliada quando um software já atendeu a aprovação funcional e de testes de unidades.