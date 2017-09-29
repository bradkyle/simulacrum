package engine

import (
	"github.com/thorad/simulacrum/asset"
	"time"
	"github.com/thorad/simulacrum/account"
)

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
	Symbol				string

	Amount				float64
	Rate				float64

	Pair				asset.Pair
	Base				asset.Asset
	Quote				asset.Asset

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

	Account				account.Account
}

func (o *Order) Side() string{
	if o.Side == Buy{
		return "Buy";
	}
	return "Sell"
}

type AnonymizedOrder struct {
	Bid         float64 `json:"bid"`
	Ask         float64 `json:"ask"`
	Shares      int     `json:"shares"`
	Ticker      string  `json:"ticker"`
	Timecreated int64   `json:"timecreated"` // unix time
}

func (o *Order) index() string {
	o.Index = o.TimeCreated.String() + o.Account.Name
	return o.Index
}


func (o *Order) price() float64 {
	K := o.Kind
	I := o.Side

	if I == Buy && K == Limit {
		return o.Bid

	} else if I == Sell && K == Limit {
		return o.Ask

	} else if I == Buy && K == Market {
		return 1000000.00

	} else if I == Sell && K == Market {
		return 0.00
	}
	// should never get here
	return 1000000.00
}
