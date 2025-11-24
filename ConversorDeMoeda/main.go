package main

import (
	"fmt"
	conversor "modConversor/Conversor"
)

func main() {
	clear()

	var origem, destino string
	var valor float64
	fmt.Print("Bem vindo ao Conversor de Moeda\n\nBRL - USD - EUR\n\nDigite o valor na moeda de origem: ")
	fmt.Scan(&valor)
	fmt.Print("Digite a moeda de origem: ")
	fmt.Scan(&origem)
	fmt.Print("Digite a moeda de destino: ")
	fmt.Scan(&destino)

	valorDoDestino := conversor.LogDoConversor(valor, origem, destino)

	clear()

	fmt.Printf("%v: %v\n%v: %v", origem, valor, destino, valorDoDestino)
}

// limpar o terminal
func clear() {
	fmt.Print("\033[H\033[2J")
}
