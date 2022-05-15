package contas

import "github.com/EnricoPDG/GolangOOP/clientes"

type ContaPoupanca struct {
	Titular                               clientes.Titular
	NumeroAgenciam, NumeroConta, Operacao int
	saldo                                 float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "O saque foi realizado com sucesso"
	}

	return "Você não tem saldo para sacar"
}

func (c *ContaPoupanca) Depositar(valorDepositar float64) (string, float64) {
	if valorDepositar > 0 {
		c.saldo += valorDepositar
		return "Deposito realizado com sucesso", c.saldo
	}
	return "Deposito não realizado", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}