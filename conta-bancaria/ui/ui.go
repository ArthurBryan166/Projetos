package ui

import (
	"fmt"
	"github.com/ArthurBryan166/Projetos/conta-bancaria/terminal"
	"github.com/ArthurBryan166/Projetos/conta-bancaria/conta"
)

func InterfaceBanco(){
	sair := false
	for !sair{
		var digiteAqui int
		fmt.Print("Bem vindo ao Tutu's Bank!!\n\nVocê deseja:\n\n1 - Acessar a conta\n2 - Criar uma conta\n0 - Sair\n\nDigite aqui: ")
		fmt.Scan(&digiteAqui)

		terminal.Clear()

		switch digiteAqui{
		case 0:
			sair = true
			fmt.Print("Muito obrigado por usar o Tutu's Bank\n")
			terminal.Pause()
			terminal.Clear()
		case 1:
			for {
				var id, senha int
				fmt.Print("Digite seu id: ")
				fmt.Scan(&id)
				fmt.Print("Digite sua senha: ")
				fmt.Scan(&senha)

				terminal.Clear()

				if id == conta.BancoGlobal.Users[id].ID && senha == conta.BancoGlobal.Users[id].Senha{
					sair3 := false
					for !sair3{
						fmt.Print("Bem vindo ")
						for _, v := range conta.BancoGlobal.Users[id].Nome{
							if v != ' '{
								fmt.Print(string(v))
							}
						}
						fmt.Print("\n\nVocê deseja:\n\n1 - Ver o seu saldo\n2 - Sacar\n3 - Depositar\n\nDigite aqui: ")
						fmt.Scan(&digiteAqui)

						terminal.Clear()

						switch digiteAqui{
						case 1:
							conta.BancoGlobal.MostrarSaldo(id)
							terminal.Pause()
							terminal.Clear()
						case 2:
							var saque float64
							fmt.Print("Digite o valor de saque: ")
							fmt.Scan(&saque)
							conta.BancoGlobal.Sacar(id, saque)
						case 3:
							var deposito float64
							fmt.Print("Digite o valor de depósito: ")
							fmt.Scan(&deposito)
							conta.BancoGlobal.Depositar(id, deposito)
						case 0:
							sair3 = true
						}
					}
				} else{
					fmt.Print("Credenciais incorretas\n")
					terminal.Pause()
					continue
				}
				break
			}
		case 2:
			var senha int
			var nome string

			fmt.Print("Digite sua senha (3 caracteres e apenas números): ")
			fmt.Scan(&senha)
			fmt.Print("Digite seu nome completo: ")
			fmt.Scan(&nome)
			terminal.Pause()
			fmt.Printf("\nSeu id é: %d\n", conta.BancoGlobal.GerarID())
			conta.BancoGlobal.CriarConta(senha, nome)
			terminal.Pause()
			terminal.Clear()
		}
	}
}