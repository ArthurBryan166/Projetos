package logicadojogo

import "fmt"

func LogDoJogo(palavraSecreta string) {
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
		fmt.Print("\nDigite uma letra: ")
		fmt.Scan(&letra)
		clear()

		// caso houver acerto de letra, este pedaço de código adiciona a letra à palavra secreta
		for i, v := range palavraSecreta{
			if string(v) == letra{
				palavra[i] = letra
				acertouLetra = true
			}
		}

		// caso erro de letra, diminui as chances
		if !acertouLetra{
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
		fmt.Print("Parabéns, você ganhou!!\n")
		func(){
				var temp string
				fmt.Scan(&temp)
			} ()
		clear()
	} else {
		fmt.Print("Que pena, você perdeu.\n")
		func(){
				var temp string
				fmt.Scan(&temp)
			} ()
		clear()
	}
}

// limpar o terminal
func clear() {
	fmt.Print("\033[H\033[2J")
}