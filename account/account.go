package account

import "time"

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

	Activity			[]Request

	HighestEquity  			float64
	HighestBalance 			float64
	LowestEquity   			float64
	LowestBalance  			float64
	FeesPaid                	float64

	MarginCalled			bool
	MarginUsed			float64
	MarginAvailable			float64
	DefaultAmount			float64					// Amount defaulted on due to not being able to pay back funding

	Balances			map[Asset]Balance
}

type Session struct{

}

type RequestKind int
const (
	Post		RequestKind = iota
	Get
)

type RequestPurpose int
const(
	GetBalances	RequestPurpose = iota
)

type Request struct{
	Kind				RequestKind
	Purpose				RequestPurpose
	Time				time.Time
}

type BalanceType int
const(
	Broker		BalanceType = iota
	BrokerMargin
	Agent
	AgentMargin
)

type Balance struct{
	Asset				Asset
	Amount				float64
	AvailableAmount			float64
	HoldAmount			float64
	Type				BalanceType
}
