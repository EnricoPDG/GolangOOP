package main

import (
	"fmt"
	//"github.com/EnricoPDG/GolangOOP/clientes"
	"github.com/EnricoPDG/GolangOOP/contas"
)

func main() {
	contaNathalo := contas.ContaPoupanca{}
	contaNathalo.Depositar(100)
	PagarBoleto(&contaNathalo, 60)
	fmt.Println(contaNathalo)

	contaEnrico := contas.ContaCorrente{}
	contaEnrico.Depositar(200)
	PagarBoleto(&contaEnrico, 40)
	fmt.Println(contaEnrico)
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}
