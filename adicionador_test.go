package main

import (
	"testing"
)

func TestAdcionador(t *testing.T) {
	soma := Adicionar(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("Esperado '%d', resultado '%d'", esperado, soma)
	}
}
