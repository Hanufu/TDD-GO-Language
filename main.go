package main

import (
	"fmt"
)

const espanhol = "Espanhol"
const frances = "Francês"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo!"
	}

	return prefixoSaudacao(idioma) + nome
}
func prefixoSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

// Adciona recebe dois inteirośee retorna a soma deles
func Adicionar(x, y int) int {
	return x + y
}

func Repetir(caractere string, quantidadeDeRepeticoes int) string {
	var repeticoes string
	if quantidadeDeRepeticoes == 0 {
		quantidadeDeRepeticoes = 1
	}
	for i := 0; i < quantidadeDeRepeticoes; i++ {
		repeticoes += caractere
	}

	return repeticoes
}

func main() {
	fmt.Println(Ola("", ""))
	fmt.Println(Repetir("a", 10))
}
