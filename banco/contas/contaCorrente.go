package contas

type ContaCorrente struct {
	Titular       string
	
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo

	if podeSacar {
		c.Saldo -= valorDoSaque
		return "O saque foi realizado com sucesso"
	}

	return "Você não tem Saldo para sacar"
}

func (c *ContaCorrente) Depositar(valorDepositar float64) (string, float64) {
	if valorDepositar > 0 {
		c.Saldo += valorDepositar
		return "Deposito realizado com sucesso", c.Saldo
	}
	return "Deposito não realizado", c.Saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}
	return false
}
