package main

import "fmt"

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

func Soma(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func SomaTudo(numerosParaSomar ...[]int) []int {
	var somas []int

	for _, numeros := range numerosParaSomar {
		somas = append(somas, Soma(numeros))
	}
	return somas
}
func main() {
	x := []int{1, 2, 3, 4}
	y := []int{5, 3, 4, 5}
	resultado := SomaTudo(x, y)
	fmt.Printf("%d", resultado)
}
