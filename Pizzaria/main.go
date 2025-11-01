package main

import "fmt"

func main() {
	var digiteAqui int
	var sair bool

	for !sair {
		fmt.Print("Bem vindo(a) à pizzaria!!\n\nVocê deseja:\n1 - Menu\n0 - Sair\n\nDigite aqui: ")
		fmt.Scan(&digiteAqui)
	}

	switch digiteAqui {
	case 0:
		sair = true
	case 1:

	}
}
