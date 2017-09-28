package app

import (
	"net/http"
	acc "github.com/thorad/simulacrum/account"
)


func (a *App) auth(r *http.Request) *acc.Account{
	r.Header.Get("X-APIKEY")
	r.Header.Get("X-PAYLOAD")	//The base64-encoded JSON payload
	r.Header.Get("X-SIGNATURE")	//hex(HMAC_SHA384(base64(payload), key=api_secret))
	return
}

func (a *App) AccountHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(GetAccountParser)
	parser = a.Parse(parser,w,r)
	response := account.GetAccount(parser)
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) AccountsHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	response := account.GetAccounts()
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) BalancesHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	response := account.GetBalances()
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) BalanceHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(GetBalanceParser)
	parser = a.Parse(parser,w,r)
	response := account.GetBalance(parser)
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) GetWithdrawalHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	response := account.GetWithdrawals()
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(WithdrawParser)
	parser = a.Parse(parser,w,r)
	response := account.Withdraw(parser)
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) GetTransferalsHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	response := account.GetTransferals()
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) TransferHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(TransferParser)
	parser = a.Parse(parser,w,r)
	response := account.Transfer(parser)
	a.Respond(w,response)
	a.spot(r)
}

