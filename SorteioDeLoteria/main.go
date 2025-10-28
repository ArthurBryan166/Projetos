package main

import (
	"fmt"
	sorteio "loteria/Sorteio"
)

func main() {
	clear()

	var quantNums int
	var numMax int

	fmt.Print("Bem vindo à loteria\n\nDigite a quantidade de números que você precisa: ")
	fmt.Scan(&quantNums)
	fmt.Print("Digite o numeral máximo: ")
	fmt.Scan(&numMax)

	clear()

	numerosAleatorios := sorteio.LogDoSorteio(quantNums, numMax)
	fmt.Println("Seus números:")
	for _, valor := range numerosAleatorios {
		fmt.Print(valor, " ")
	}
}

// limpar terminal
func clear() {
	fmt.Print("\033[H\033[2J")
}
