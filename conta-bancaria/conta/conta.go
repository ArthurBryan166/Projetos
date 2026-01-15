package conta

type ContaBancaria struct{
	ID int
	Senha int
	Nome string
	Saldo float64
}

type Banco struct{
	Users map[int]*ContaBancaria
}

func inicializarBanco() *Banco{
	return &Banco{
		Users: make(map[int]*ContaBancaria),
	}
}

var BancoGlobal = inicializarBanco()

func (u *Banco) CriarConta(password int, name string){
	id := u.GerarID()
	u.Users[id] = &ContaBancaria{
		ID: id,
		Senha: password,
		Nome: name,
	}
}