package main

import (
	"fmt"
	"loteria/sorteio"
)

func main() {
	sair := false
	for !sair {
		clear()

		var quantNums int
		var numMax int
		var sairDigiteAqui int

		fmt.Print("Bem vindo à loteria\n\n1 - Loteria\n0 - Sair\n\nDigite aqui: ")
		fmt.Scan(&sairDigiteAqui)

		clear()

		switch sairDigiteAqui {
		case 0:
			sair = true
		case 1:
			fmt.Print("Digite a quantidade de números que você precisa: ")
			fmt.Scan(&quantNums)
			fmt.Print("Digite o numeral máximo: ")
			fmt.Scan(&numMax)

			clear()

			numerosAleatorios := sorteio.LogDoSorteio(quantNums, numMax)
			fmt.Println("Seus números:")
			for _, valor := range numerosAleatorios {
				fmt.Print(valor, " ")
			}
			func() {
				temp := ""
				fmt.Scan(&temp)
			}()
		}
	}
}

// limpar terminal
func clear() {
	fmt.Print("\033[H\033[2J")
}
