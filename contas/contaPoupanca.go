package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operação      int
	saldo         float64
}

func (c *ContaPoupanca) Transferir(valorTransferencia float64, contaDestino *ContaPoupanca) bool {
	if valorTransferencia <= c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Deposito(valorTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso, saldo atual: ", c.saldo
	} else {
		return "Saldo insuficiente: ", c.saldo
	}
}

func (c *ContaPoupanca) Deposito(valorDeposito float64) (string, float64) {
	podeDeposito := valorDeposito > 0

	if podeDeposito {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso: ", c.saldo
	} else {
		return "Deposito errado: ", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
