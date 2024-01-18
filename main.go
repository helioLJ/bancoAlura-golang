package main

import (
	"fmt"
	"bancoAlura/contas"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	fmt.Println(contaDaSilvia.Saldo)
	fmt.Println(contaDoGustavo.Saldo)

	fmt.Println(contaDaSilvia.Transferir(300, &contaDoGustavo))

	fmt.Println(contaDaSilvia.Saldo)
	fmt.Println(contaDoGustavo.Saldo)
}