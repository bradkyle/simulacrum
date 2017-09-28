package app

import (
	"github.com/thorad/simulacrum/account"
	"net/http"
)

type Agent struct{
	Scope 		string
}

func (a App) NewAccount(parser NewAccountParser) (*account.Account, error){

	return &account.Account{}, nil
}

func (a *App) auth(r *http.Request) *account.Account{

	return
}

func (a *App) spot(r *http.Request){

}

func (a *App) rule(w http.ResponseWriter,r *http.Request){

}

