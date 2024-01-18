package main

import (
	"bancoAlura/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64)  {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string  
}

func main() {
	contaDoHelio := contas.ContaPoupanca{}
	contaDoHelio.Depositar(900)
	PagarBoleto(&contaDoHelio, 60)

	fmt.Println(contaDoHelio.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(600)

	PagarBoleto(&contaDaLuisa, 400)

	fmt.Println(contaDaLuisa.ObterSaldo())
}
