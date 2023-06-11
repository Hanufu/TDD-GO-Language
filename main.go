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

func main() {
	fmt.Print(Ola("", ""))
}
