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

	for len(nums) < int(quantNumsAleatorios) {

	}
}
