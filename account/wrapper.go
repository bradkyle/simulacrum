package account

import "github.com/thorad/simulacrum/parse"

func NewAccount(parser parse.NewAccountParser) parse.NewAccountParser{

}

func (a *Account) GetAccount(parser parse.GetAccountParser) parse.GetAccountParser{

}

func (a *Account) GetAccounts() parse.GetAccountsParser{

}

func (a *Account) GetBalances() parse.GetBalancesParser{

}

func (a *Account) GetBalance(parser  parse.GetBalanceParser) parse.GetBalanceParser{

}

func (a *Account) GetWithdrawals()  parse.GetWithdrawalsParser{

}

func (a *Account) Withdraw(parser parse.WithdrawParser) parse.WithdrawParser{

}

func (a *Account) GetTransferals()  parse.GetTransferalsParser{

}

func (a *Account) Transfer(parser parse.TransferParser) parse.TransferParser{

}