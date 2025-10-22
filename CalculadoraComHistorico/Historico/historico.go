package historico

import "fmt"

type Historico struct {
	Num1      float64
	Num2      float64
	Resultado float64
	Operador  string
}

var operacoes []Historico

func HistoricoCalc(n1 float64, n2 float64, res float64, opr string) {
	operacoes = append(operacoes, Historico{Num1: n1, Num2: n2, Resultado: res, Operador: opr})
}

func MostrarHistorico() {
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
