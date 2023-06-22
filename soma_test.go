package main

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {
	t.Run("Coleção de 5 números", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}
		resultado := Soma(numeros)
		esperado := 15

		if esperado != resultado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}
	})

	t.Run("Coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}

		resultado := Soma(numeros)
		esperado := 6

		if resultado != esperado {
			t.Errorf("Resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}
	})
}

func TestSomaTudo(t *testing.T) {

	resultado := SomaTudo([]int{1, 2}, []int{0, 9})

	esperado := []int{3, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("Resultado %v, esperado %v", resultado, esperado)
	}
}
