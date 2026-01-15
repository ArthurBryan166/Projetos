package conta

func (u *Banco) Sacar(id int, valor float64){
	u.Users[id].Saldo -= valor
}