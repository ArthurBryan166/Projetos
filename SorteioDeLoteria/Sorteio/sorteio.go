package sorteio

import (
	"math/rand"
	"time"
)

func LogDoSorteio(quantNumsAleatorios, numeralMax uint) []uint {
	// numero aleatorio
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	var nums []uint

	for len(nums) != int(quantNumsAleatorios) {
		/*  a atribuicao de r nao é assim, revisar a forma correta
			nao esquecer do r % numeralMax + 1
		*/
		numAleatorio := uint(r)
		existe := false

		for _, valor := range nums{
			if numAleatorio == valor{
				existe = true
			}
		}

		if !existe{
			nums = append(nums, numAleatorio)
		}
	}
}
