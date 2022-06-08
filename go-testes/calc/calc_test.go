package calc

// O pacote testing é importado
import "testing"

func TestSum(t *testing.T) {
	// Os dados a serem usados no teste são inicializados (entrada/saída)
	num1 := 3
	num2 := 5
	expectedResult := 8

	// O teste é executado
	result := Sum(num1, num2)

	// Os resultados são validados
	if result != expectedResult {
		t.Errorf("A função Sum() retornou o resultado = %v, mas o esperado é %v", result, expectedResult)
	}
}
