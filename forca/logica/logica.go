package logica

import (
	"fmt"
	"modJogo/terminal"
	"modJogo/historico"
)

func UI(palavraSecreta string){
	// variáveis
	var palavra []string
	var letra string
	chances := 10
	var acertou, acertouLetra, ganhou bool

	// palavra secreta que aparecerá no terminal
	for i := 0; i < len(palavraSecreta); i++ {
		palavra = append(palavra, "-")
	}

	// loop do jogo
	for chances > 0{
		// variáveis
		acertou = true
		acertouLetra = false

		// tela do jogo
		fmt.Printf("Chances: %d\nPalavra:\n", chances)
		for _, letras := range palavra{
			fmt.Print(letras)
		}
		fmt.Print("\nLetras já digitadas:")
		historico.MostrarLetras()
		fmt.Print("\n\nDigite uma letra: ")
		fmt.Scan(&letra)
		terminal.Clear()

		// caso houver acerto de letra, este pedaço de código adiciona a letra à palavra secreta
		for i, v := range palavraSecreta{
			if string(v) == letra{
				palavra[i] = letra
				acertouLetra = true
			}
		}

		// caso erro de letra, diminui as chances
		if !acertouLetra && !historico.AddLetras(letra){
			chances--
		}

		// verificar se o jogador ganhou o jogo
		for _, v := range palavra{
			if v == "-"{
				acertou = false
				break
			}
		}

		// break do loop caso haja vitória
		if acertou{
			chances = 0
			ganhou = true
		}
	}

	// mensagens de vitória e derrota
	if ganhou{
		fmt.Printf("Parabéns, você ganhou!! A palavra era %s\n", palavraSecreta)
		terminal.Pause()
		terminal.Clear()
	} else {
		fmt.Printf("Que pena, você perdeu. A palavra era %s\n", palavraSecreta)
		terminal.Pause()
		terminal.Clear()
	}
}