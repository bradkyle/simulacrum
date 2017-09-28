package app

import "net/http"

func (a *App) AccountHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(GetAccountParser)
	parser = a.Parse(parser,w,r)
	response, err := account.GetAccount(parser)
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) AccountsHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	response, err := account.GetAccounts()
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) BalancesHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	response, err := account.GetBalances()
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) BalanceHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(GetBalanceParser)
	parser = a.Parse(parser,w,r)
	response, err := account.GetBalance()
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) GetWithdrawalHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	response, err := account.GetWithdrawals()
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(WithdrawParser)
	parser = a.Parse(parser,w,r)
	response, err := account.Withdraw(parser)
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) GetTransferalsHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	response, err := account.GetTransferals()
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) TransferHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(TransferParser)
	parser = a.Parse(parser,w,r)
	response, err := account.Transfer(parser)
	a.Respond(w,response,err)
	a.spot(r)
}

