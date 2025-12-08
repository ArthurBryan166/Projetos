package main

import (
	"fmt"
	"modJogo/LogicaDoJogo"
)

func main() {
	var sair bool
	for !sair {
		var digiteAqui int
		fmt.Print("Bem vindo ao Jogo da Forca\n\nVocê deseja:\n\n1 - Jogar\n0 - Sair\n\nDigite aqui: ")
		fmt.Scan(&digiteAqui)
		switch digiteAqui{
		case 0:
			sair = true
			clear()
			fmt.Print("Obrigado por jogar o jogo da forca\n\n")
		case 1:
			var palavra string
			fmt.Print("Digite a palavra secreta (apenas letras, e letras sem acento): ")
			fmt.Scan(&palavra)
			logicadojogo.LogDoJogo(palavra)
		default:
			fmt.Print("Digite um número válido\n")
			func(){
				var temp string
				fmt.Scan(&temp)
			} ()
			clear()
		}
	}
}

// limpar o terminal
func clear() {
	fmt.Print("\033[H\033[2J")
}