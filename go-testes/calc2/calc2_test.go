package calc2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	// Os dados a serem usados no teste são inicializados (entrada/saída)
	num1 := 3
	num2 := 5
	expectedResult := 8

	// O teste é executado
	result := Sum(num1, num2)

	// Os resultados são validados
	assert.Equal(t, expectedResult, result, "devem ser iguais")
}

func TestDivide(t *testing.T) {
	// Os dados a serem usados no teste são inicializados (entrada/saída)
	num1 := 3
	num2 := 0

	// O teste é executado
	result := Divide(num1, num2)

	// Os resultados são validados
	assert.NotNil(t, result)
}
