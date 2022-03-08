package hello

import "testing"

func testOla(t *testing.T) {
	resultado := Ola("mundo")
	esperado := "Olá, mundo!"

	if resultado != esperado {
		t.Errorf("Resultado '%s', esperado '%s'", resultado, esperado)
	}
}
