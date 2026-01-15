package conta

import "fmt"

func (u Banco) MostrarSaldo(id int) {
	fmt.Printf("Saldo: %.2f\n", u.Users[id].Saldo)
}

func (u *Banco) Depositar(id int, valor float64){
	u.Users[id].Saldo += valor
}