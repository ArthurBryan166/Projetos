package calculadora

import (
	historico "calc/Historico"
	"fmt"
)

func Calculadora() {
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

	historico.HistoricoCalc(num1, num2, resultado, operador)

	fmt.Println("\nResultado:", resultado)
	fmt.Scan(&temp)
	clear()
}

func clear() {
	fmt.Print("\033[H\033[2J")
}
