package account

import (
	"github.com/thorad/simulacrum/asset"
	"github.com/thorad/simulacrum/engine"
	"errors"
)

type AccountStatus int
const(
	Pending		AccountStatus = iota
	Setup
	Blocked
)
type Account struct{
	Index				string
	Name                            string
	Status				AccountStatus
	HighestEquity  			float64
	HighestBalance 			float64
	LowestEquity   			float64
	LowestBalance  			float64
	FeesPaid                	float64
	MarginCalled			bool
	MarginUsed			float64
	MarginAvailable			float64
	DefaultAmount			float64					// Amount defaulted on due to not being able to pay back funding
	Balances			map[asset.Asset]Balance
}

type BalanceType int
const(
	Broker		BalanceType = iota
	BrokerMargin
	Agent
	AgentMargin
)

type Balance struct{
	Asset				asset.Asset
	Amount				float64
	AvailableAmount			float64
	HoldAmount			float64
	Type				BalanceType
}


func (a *Account) buy(t engine.Trade) error{
	if a.Balances[t.Base] >= (t.Amount * t.Price) {
		a.Balances[t.Base] = a.Balances[t.Base] - (t.Amount * t.Price)
		a.Balances[t.Target] = a.Balances[t.Target] + t.Amount
	}
	return errors.New("")
}

func (a *Account) sell(t engine.Trade) {
	if a.Balances[t.Target] >= (t.Amount * t.Price) {
		a.Balances[t.Base] = a.Balances[t.Base] + (t.Amount * t.Price)
		a.Balances[t.Target] = a.Balances[t.Target] + t.Amount
	}
}


