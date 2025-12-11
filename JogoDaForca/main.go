package main

import (
	"fmt"
	"modJogo/logicadojogo"
	"modJogo/terminal"
)

func main() {
	var sair bool
	for !sair {
		var digiteAqui int
		terminal.Clear()
		fmt.Print("Bem vindo ao Jogo da Forca\n\nVocê deseja:\n\n1 - Jogar\n0 - Sair\n\nDigite aqui: ")
		fmt.Scan(&digiteAqui)
		terminal.Clear()
		switch digiteAqui{
		case 0:
			sair = true
			terminal.Clear()
			fmt.Print("Obrigado por jogar o jogo da forca\n\n")
		case 1:
			var palavra string
			fmt.Print("Digite a palavra secreta (apenas letras, e letras sem acento): ")
			fmt.Scan(&palavra)
			terminal.Clear()
			logicadojogo.UI(palavra)
		default:
			fmt.Print("Digite um número válido\n")
			terminal.Pause()
			terminal.Clear()
		}
	}
}