package logicadoprograma

import (
	"fmt"
	"math/rand"
	"time"
)

func LogDoJogo(chances int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numAleatorio := r.Intn(int(100) + 1)

	for chances > 0 {
		var numDigitado int
		fmt.Printf("Chances: %d\nDigite um número: ", chances)
		fmt.Scan(&numDigitado)
		clear()

		if numDigitado == numAleatorio {
			fmt.Printf("Parabéns, você adivinhou o número secreto!!\n")
			func (){
				var temp rune
				fmt.Scan(&temp)
			} ()
			break
		} else if numDigitado < numAleatorio {
			fmt.Printf("O número aleatório é maior que %d\n", numDigitado)
		} else {
			fmt.Printf("O número aleatório é menor que %d\n", numDigitado)
		}
		func (){
			var temp rune
			fmt.Scan(&temp)
		} ()
		clear()
		chances--
	}
	if chances == 0 {
		fmt.Print("Que pena, você perdeu!!")
	}
}

// limpar o terminal
func clear() {
	fmt.Print("\033[H\033[2J")
}
