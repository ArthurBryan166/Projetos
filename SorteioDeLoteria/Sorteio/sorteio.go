package sorteio

import (
	"math/rand"
	"time"
)

func LogDoSorteio(quantNumsAleatorios, numeralMax int) []int {
	// numero aleatorio
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var nums []int
	var numAleatorio int

	for len(nums) < quantNumsAleatorios {
		numAleatorio = r.Intn(int(numeralMax) + 1)
		existe := false

		for _, valor := range nums {
			if numAleatorio == valor {
				existe = true
			}
		}

		if !existe {
			nums = append(nums, numAleatorio)
		}
	}
	return nums
}
