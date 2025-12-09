package logicadojogo

import "fmt"

func LogDoJogo(palavraSecreta string) {
	var palavra []string
	var letra string
	chances := 10
	var acertou, acertouLetra bool

	for i := 0; i < len(palavraSecreta); i++ {
		palavra = append(palavra, "-")
	}

	for chances > 0{
		acertou = true
		acertouLetra = false
		fmt.Printf("Chances: %d\n\nPalavra:\n", chances)
		for _, letras := range palavra{
			fmt.Print(letras)
		}
		fmt.Print("Digite uma letra: ")
		fmt.Scan(&letra)

		for i, v := range palavraSecreta{
			if string(v) == letra{
				palavra[i] = letra
				acertouLetra = true
			}
		}

		if !acertouLetra{
			chances--
		}

		for _, v := range palavra{
			if v == "-"{
				acertou = false
				break
			}
		}

		if acertou{
			chances = 0
		}
	}
}