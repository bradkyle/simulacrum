package app

import "net/http"

func (a *App) AccountHandler(w http.ResponseWriter, r *http.Request) {
	account := a.auth(r)
	parser := new(GetAccountParser)
	parser.Parse(w,r)
	response, err := account.GetAccount(parser)
	write(w,response,err)
}

func (a *App) AccountsHandler(w http.ResponseWriter, r *http.Request) {
	account := a.auth(r)
	response, err := account.GetAccounts()
	write(w,response,err)
}

func (a *App) BalancesHandler(w http.ResponseWriter, r *http.Request) {
	account := a.auth(r)
	response, err := account.GetBalances()
	write(w,response,err)
}

func (a *App) BalanceHandler(w http.ResponseWriter, r *http.Request) {
	account := a.auth(r)
	parser := new(GetBalanceParser)
	parser.Parse(w,r)
	response, err := account.GetBalance()
	write(w,response,err)
}

func (a *App) GetWithdrawalHandler(w http.ResponseWriter, r *http.Request) {
	account := a.auth(r)
	response, err := account.GetWithdrawals()
	write(w,response,err)
}

func (a *App) WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	account := a.auth(r)
	parser := new(WithdrawParser)
	parser.Parse(w,r)
	response, err := account.Withdraw(parser)
	write(w,response,err)
}

func (a *App) GetTransferalsHandler(w http.ResponseWriter, r *http.Request) {
	account := a.auth(r)
	response, err := account.GetTransferals()
	write(w,response,err)
}

func (a *App) TransferHandler(w http.ResponseWriter, r *http.Request) {
	account := a.auth(r)
	parser := new(TransferParser)
	parser.Parse(w,r)
	response, err := account.Transfer(parser)
	write(w,response,err)
}

