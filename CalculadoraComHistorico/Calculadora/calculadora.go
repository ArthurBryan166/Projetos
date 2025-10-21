package main

import "fmt"

type Historico struct {
	Num1      float64
	Num2      float64
	Resultado float64
	Operador  string
}

var operacoes []Historico

func main() {
	clear()
	sair := 1
	for sair != 0 {
		var digiteAqui int
		fmt.Print("Bem vindo à calculadora\n\n1 - Calculadora\n2 - Histórico\n0 - Sair\n\nDigite aqui: ")
		fmt.Scan(&digiteAqui)

		switch digiteAqui {
		case 0:
			sair = 0
			clear()
		case 1:
			calculadora()
		case 2:
			fmt.Print("Histórico:\n\n")
			mostrarHistorico()
		}
	}
}

func calculadora() {
	clear()
	var num1, num2, resultado float64
	var operador, temp string

	fmt.Print("Digite um número: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o operador: ")
	fmt.Scan(&operador)
	fmt.Print("Digite outro número: ")
	fmt.Scan(&num2)

	switch operador {
	case "+":
		resultado = num1 + num2
	case "-":
		resultado = num1 - num2
	case "*":
		resultado = num1 * num2
	case "/":
		resultado = num1 / num2
	}

	historico(num1, num2, resultado, operador)

	fmt.Println("\nResultado:", resultado)
	fmt.Scan(&temp)
	clear()
}

func historico(n1 float64, n2 float64, res float64, opr string) {
	operacoes = append(operacoes, Historico{Num1: n1, Num2: n2, Resultado: res, Operador: opr})
}

func mostrarHistorico() {
	clear()
	var temp string
	for _, op := range operacoes {
		fmt.Printf("%v %v %v = %v\n", op.Num1, op.Operador, op.Num2, op.Resultado)
	}
	fmt.Scan(&temp)
	clear()
}

func clear() {
	fmt.Print("\033[H\033[2J")
}
