package main

import (
	"fmt"
	c "mymodule/contas"
)

func main() {
	contaDoEnrico := c.ContaCorrente{}
	contaDoEnrico.Titular = "Enrico"
	contaDoEnrico.Saldo = 100

	contaDaNath := c.ContaCorrente{
		Titular: "Nathalia",
		Saldo: 001,
	}
	
	status := contaDoEnrico.Transferir(50, &contaDaNath)

	fmt.Println(contaDoEnrico)
	fmt.Println(contaDaNath)
	fmt.Println(status)
	fmt.Println(contaDoEnrico)
	fmt.Println(contaDaNath)
}