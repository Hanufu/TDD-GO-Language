package main

import "testing"

func TestReperit(t *testing.T) {
	verificarRepeticaoCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("Resultado '%s' esperado '%s'", resultado, esperado)
		}
	}
	t.Run("Faz a repeticao padrão de algum caractere", func(t *testing.T) {
		resultado := Repetir("a", 5)
		esperado := "aaaaa"
		verificarRepeticaoCorreta(t, resultado, esperado)
	})

	t.Run("Faz repeticão passando zero", func(t *testing.T) {
		retultado := Repetir("a", 0)
		esperado := "a"
		verificarRepeticaoCorreta(t, retultado, esperado)
	})
}
