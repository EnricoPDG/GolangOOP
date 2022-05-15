package contas

import "github.com/EnricoPDG/GolangOOP/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "O saque foi realizado com sucesso"
	}

	return "Você não tem saldo para sacar"
}

func (c *ContaCorrente) Depositar(valorDepositar float64) (string, float64) {
	if valorDepositar > 0 {
		c.saldo += valorDepositar
		return "Deposito realizado com sucesso", c.saldo
	}
	return "Deposito não realizado", c.saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}
	return false
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
