package historico

import "slices"

import "fmt"

var letras []string

func AddLetras(letra string) bool{
	if slices.Contains(letras, letra) {
		return true
	}
	letras = append(letras, letra)
	return false
}

func MostrarLetras() {
	for _, v := range letras {
		fmt.Printf(" %s", v)
	}
}