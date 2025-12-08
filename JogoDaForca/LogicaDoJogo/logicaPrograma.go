package logicadojogo

import "fmt"

func LogDoJogo(palavraSecreta string) {
	var palavra []string
	chances := 10

	for i := 0; i < len(palavraSecreta); i++ {
		palavra = append(palavra, "-")
	}

	fmt.Printf("Chances: %d\n\nPalavra:\n")
	for _, letras := range palavra{
		fmt.Print(letras)
	}
	
}