package engine

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