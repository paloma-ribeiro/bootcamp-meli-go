# Princípio F.I.R.S.T.

## F - Fast

É possível ter milhares de testes em seu projeto e eles devem ser rápidos de executar.

## I - Isolated/Independent

Um método de teste deve estar em conformidade com o "3A" (Arrange, Act, Assert), ou o que é o mesmo: give, when, then. Além disso, eles não precisam ser executados em uma determinada ordem para funcionar.

## R - Repeatable

Resultados determinísticos. Eles não devem depender de dados do ambiente enquanto estiverem em execução, por exemplo, a hora do sistema.

## S - Self-Validating

A inspeção manual não deve ser necessária para validar os resultados.

## T - Thorough

Eles devem cobrir todos os cenários de um caso de uso, e não apenas almejar 100% de cobertura. Mutações de teste, casos extremos, exceções, erros, etc.