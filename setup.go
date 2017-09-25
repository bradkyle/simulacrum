package main

import (
	"time"
	"github.com/gorilla/mux"
)

type App struct{
	Router				*mux.Router
	Config				Config
	ConfigFile			string
	Logger				Logger
	shutdown 			chan bool


	Accounts			map[string]*Account
	Orderbooks			map[string]*Orderbook
	Tickers				map[string]*Ticker
	Engine				Engine
}

type Logger struct{

}



type MatchingEngineType int
const(
	PriceTimePriority MatchingEngineType = iota
)

type Orderbook struct{
	Engine 				MatchingEngineType
	Pair				Pair
	BuyQueue			*Orderqueue
	SellQueue			*Orderqueue
}

type Engine interface{

}

type Orderqueue struct{

}

type Ticker struct {

}

type Tick struct{
	Interval			int64
	Symbol				string
	BaseSymbol			string
	QuoteSymbol			string
	LastPrice			float64
	MidPrice			float64					// bid + ask / 2
	LowPrice			float64					// Lowest trade price of the last 24 hours
	HighPrice			float64
	LowestAsk			float64
	HighestAsk			float64
	Volume        			float64
	VolumeWeightedAveragePrice	float64
	NumberOfTrades			float64
	DayOpeningPrice			float64
	ValidTime			float64
	VolumeBase			float64
	VolumeQuote			float64
}

type FundingBookType int

type Fundingbook struct{

}

type Funding struct{

}

type Pair struct{
	Id				int64
	Symbol				string
	Base				Asset
	Quote				Asset
	MinimumOrderSize		float64
	MaximumOrderSize		float64
	MinimumOrderIncrement		float64
	MinimumPriceIncrement		float64
	MinimumMargin			float64
	InitMargin			float64
	Enabled				bool
}

type AssetType int
const(
	Cryptocurrency	AssetType = iota
	Fiat
	Bond
	Commodity
	Share
	CTFO
)

type Asset struct{
	Name				string
	Symbol				string
	AssetType			AssetType
	Enabled				bool
}




type OrderKind int
const(
	Limit	OrderKind = iota
	Market
)

type OrderSide int
const(
	Buy	OrderSide = iota
	Sell
)

type OrderExecution int
const(
	MarginMarket	OrderExecution = iota
	MarginLimit
	MarginStop
	MarginTrailingStop
	MarginFillOrKill
	ExchangeMarket
	ExchangeLimit
	ExchangeStop
	ExchangeTrailingStop
	ExchangeFillOrKill
	BrokerMarket
	BrokerLimit
	BrokerStop
	BrokerTrailingStop
	BrokerFillOrKill
)

type Order struct{
	Index				string

	Amount				float64
	Rate				float64

	Pair				Pair
	Base				Asset
	Quote				Asset

	Kind				OrderKind
	Side				OrderSide
	Execution			OrderExecution

	Active				bool
	Closed				bool
	Filled				bool
	Expired				bool

	ExpiresAt			time.Time
	TimeCreated			time.Time
	TimeClosed			time.Time
	TimeFilled			time.Time
	TimeStopped			time.Time

	StopLoss			float64
	TakeProfit			float64
	TrailingStop			float64

	TakeProfitHit			bool
	StopLossHit			bool
	TrailingStopHit 		bool

	OpenBid  			float64
	OpenAsk  			float64
	CloseBid 			float64
	CloseAsk 			float64

	EquityAtOpen   			float64
	EquityAtClose  			float64
	BalanceAtOpen  			float64
	BalanceAtClose 			float64
	LowestBid      			float64
	LowestAsk      			float64
	HighestBid     			float64
	HighestAsk     			float64
	OrdersOpenAtOpen  		int64
	OrdersOpenAtClose 		int64
	DrawdownAtOpen  		float64
	DrawdownAtClose 		float64
	ClosestTakeProfit 	  	float64
	ClosestStopLoss      		float64
	Delay                   	time.Duration

	Account				Account
}


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

type Wallet struct{
	Address				string
	Asset				Asset
	Amount				float64
}

type Withdrawal struct{
	Address				string
	Asset				Asset
	Amount				float64
}





