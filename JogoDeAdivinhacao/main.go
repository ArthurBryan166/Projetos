package main

import (
	"fmt"
	logicadoprograma "modJogo/LogicaDoPrograma"
)

func main() {
	sair := false
	for !sair {
		var digiteAqui int
		fmt.Print("Bem vindo ao jogo de adivinhação!!\n\nVocê deseja jogar no nível\n\n1 - Fácil (10 chances)\n2 - Médio (7 chances)\n3 - Difícil (5 chances)\n0 - Sair\n\nDigite aqui: ")
		fmt.Scan(&digiteAqui)
		switch digiteAqui {
		case 0:
			sair = true
		case 1:
			logicadoprograma.LogDoJogo(10)
		case 2:
			logicadoprograma.LogDoJogo(7)
		case 3:
			logicadoprograma.LogDoJogo(5)
		}
	}
}
