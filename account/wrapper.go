package account

import "github.com/thorad/simulacrum/app"

func NewAccount(parser app.NewAccountParser) app.NewAccountParser{

}

func (a *Account) GetAccount(parser app.GetAccountParser) app.GetAccountParser{

}

func (a *Account) GetAccounts() app.GetAccountsParser{

}

func (a *Account) GetBalances() app.GetBalancesParser{

}

func (a *Account) GetBalance(parser  app.GetBalanceParser) app.GetBalanceParser{

}

func (a *Account) GetWithdrawals()  app.GetWithdrawalsParser{

}

func (a *Account) Withdraw(parser app.WithdrawParser) app.WithdrawParser{

}

func (a *Account) GetTransferals()  app.GetTransferalsParser{

}

func (a *Account) Transfer(parser app.TransferParser) app.TransferParser{

}