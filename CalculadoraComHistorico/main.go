package main

import (
	calculadora "calc/Calculadora"
	historico "calc/Historico"
	"fmt"
)

func main() {
	clear()
	sair := 1

	// looping de uso da calculadora
	for sair != 0 {
		var digiteAqui int
		fmt.Print("Bem vindo à calculadora\n\n1 - Calculadora\n2 - Histórico\n0 - Sair\n\nDigite aqui: ")
		fmt.Scan(&digiteAqui)

		switch digiteAqui {
		case 0:
			sair = 0
			clear()
			fmt.Println("Obrigado por usar a calculadora!!")
		case 1:
			calculadora.Calculadora()
		case 2:
			fmt.Print("Histórico:\n\n")
			historico.MostrarHistorico()
		}
	}
}

// limpar o terminal
func clear() {
	fmt.Print("\033[H\033[2J")
}
