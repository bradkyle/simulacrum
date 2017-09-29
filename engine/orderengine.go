package engine

import (
	"github.com/thorad/simulacrum/account"
	"time"
)

// Continuous Trading	----------------------------------------------------------------------------------------------->

/*
FifoEngine: Price/Time Algorithm
The FIFO algorithm uses price and time as the only criteria for filling an order. In this algorithm,
all orders at the same price level are filled according to time priority; the first order at a price
level is the first order matched.

Using the FIFO match algorithm:
Orders at the same price are prioritized by their entry times, with the oldest order having the highest priority.
The aggregated implied order has the lowest priority within its price level.
In the example, if any orders on the offer side were to come in at the corresponding price of 9330,
then the 3-lot and 5-lot would match first (in that order), with the 2-lot implied order filled last
because it has the lowest priority.
*/
type FifoEngine struct{
	OrderEngine
}

func (e *FifoEngine) run(symbol string) {
		buyTop := e.queue.Peek("BUY" + symbol)
		sellTop := e.queue.Peek("SELL" + symbol)

		for buyTop != nil && sellTop != nil && buyTop.Value >= sellTop.Value {
			buy := e.store.get(buyTop.Lookup)
			sell := e.store.get(sellTop.Lookup)
			price := e.negotiate(buy, sell)

			if buy.Amount == sell.Amount {
				e.queue.Dequeue("BUY" + symbol)
				e.queue.Dequeue("SELL" + symbol)

				go e.trade(buy.fill(price), sell.fill(price))

				e.store.remove(buyTop.Lookup)
				e.store.remove(sellTop.Lookup)

			} else if buy.Amount < sell.Amount {
				e.queue.Dequeue("BUY" + symbol)
				remainderSell := sell.Amount - buy.Amount

				go e.trade(buy.fill(price), sell.partialFill(price, remainderSell))

				e.store.remove(buyTop.Lookup)

			} else if buy.Amount > sell.Amount {
				e.queue.Dequeue("SELL" + symbol)
				remainderBuy := buy.Amount - sell.Amount

				go e.trade(sell.fill(price), sell.partialFill(price, remainderBuy))

				e.store.remove(sellTop.Lookup)
			}

			buyTop = e.queue.Peek("BUY" + symbol)
			sellTop = e.queue.Peek("SELL" + symbol)
		}
		e.Publish(symbol, buyTop, sellTop)
}



/*
FifoLMMEngine: Price/Time Algorithm
The FIFO with LMM algorithm is an enhanced FIFO algorithm that allows for LMM allocations prior to the
FIFO allocations. LMMs are each allocated a configurable percentage of an aggressor order before the
remaining quantity is allocated FIFO.
The FIFO with LMM algorithm follows these stages to match trades:
	- LMM set to configurable percent
	- FIFO allocation
*/

type FifoLMMEngine struct{
	OrderEngine
}

func (e *FifoLMMEngine) run(symbol string) {

}



/*
FifoTopLMMEngine: Price/Time Algorithm
The FIFO with LMM algorithm is an enhanced FIFO algorithm that allows for LMM allocations prior
to the FIFO allocations. Additionally, this algorithm incorporates a priority (top order) to the
first incoming order that betters the market. If priority is established, the aggressor orders are
first allocated to the top order until the order's quantity is exhausted.
The FIFO with LMM algorithm follows these stages to match trades:
	- Assigns a top order percent allocation that betters the market
	- LMM set to configurable percent
	- FIFO allocation
*/

type FifoTopLMMEngine struct{
	OrderEngine
}

func (e *FifoTopLMMEngine) run(symbol string) {

}


/*
FifoOneTradeTopLMMEngine: FIFO/ProRata Algorithm
The Split FIFO/Pro-Rata algorithm is a hybrid which integrates a percent-based allocation on both a
FIFO and Pro-Rata formula to the resting order book. This algorithm has been developed with the flexibility
to calibrate the level of tradable quantity that is allocated on a FIFO and pro-rata basis (X% FIFO,
Y% pro-rata for every tradable quantity).
Additionally, an order that does not receive an allocation during the Pro-Rata stage is eligible for a one
lot fill through pro-rata leveling. The allocations for this component are prioritized for orders with the
largest remaining quantity. If multiple orders have the same remaining quantity time priority will be the
determining factor.
This algorithm can also incorporate a top order and LMM components before the FIFO and pro-rata
allocations are calculated.
The Split FIFO/Pro-Rata algorithm follows these stages to match trades:
	- Assigns a top order percent allocation (with min and/or max)
	- LMM may be applied
	- FIFO percent may be configured from 0 to 100%
	- Pro-rata with min may be configured from 0 to 100%
	- Pro-rata leveling
	- FIFO for any residual quantity
*/

type FifoProRataEngine struct{
	OrderEngine
}

func (e *FifoProRataEngine) run(symbol string) {

}



/*
ProRataEngine: Pro-Rata Algorithm
Under this algorithm, an incoming trade is split among limit orders proportionally to their
sizes.
The FX Calendar algorithm fills orders according to price, order lot size and time.
An incoming aggressor order's quantity is multiplied by each resting order's pro-rated percentage
to calculate allocated trade quantity. All fills are rounded down to the nearest integer; if an
allocated trade quantity is less than two, it is rounded down to zero.
An order's pro-rata percentage is calculated by taking order quantity divided by total quantity at
a certain price. Excess lots, which occur as a result of the rounding down of the original allocated
trade quantity, may be allocated FIFO.
The FX Calendar algorithm follows these stages to match trades:
	- Pro-rata with a minimum allocation
	- FIFO for any residual quantity
*/

type ProRataEngine struct{
	OrderEngine
}

func (e *ProRataEngine) run(symbol string) {
	
}



/*
ThresholdProRataEngine: Pro-Rata Algorithm
Under this algorithm, an incoming trade is split among limit orders proportionally to their
sizes.
The Threshold Pro-Rata algorithm is an enhanced pro-rata algorithm that incorporates a priority (top order)
to the first incoming order that betters the market. If priority is established, aggressor orders are first
allocated to the top order up to the maximum configured priority allocation (volume cap). A minimum order
size may also be set, whereas if an order does not meet that minimum level, it does not qualify for top order.
Additionally, there is a minimum configurable allocation parameter (pro-rata mini), which is set to a quantity
that can be allocated due to the results from the initial pro-rated allocation. All fills are rounded down to
the nearest integer and if an allocated trade quantity is less than the pro-rata min it is rounded down to zero.
Excess lots are allocated FIFO.
The Threshold Pro-Rata algorithm follows these stages to match trades:
	- Assigns a top order percent allocation (with min and/or max)
	- Pro-rata with configurable minimum allocation
	- FIFO for any residual quantity
*/

type ThresholdProRataEngine struct{
	OrderEngine
}

func (e *ThresholdProRataEngine) run(symbol string) {

}


/*
ThresholdProRataLMMEngine: Pro-Rata Algorithm (Market Maker)
Under this algorithm, an incoming trade is split among limit orders proportionally to their
sizes.
The Threshold Pro-Rata with LMM algorithm is an enhanced pro-rata algorithm that incorporates a Lead Market
Maker (LMM) allocation after top order and before the pro-rata allocation.
LMM allows CME Group-designated customers to receive a designated allocation percentage of aggressor
orders. The allocated quantity will not exceed the incoming order's quantity. Multiple LMM orders are matched
based on FIFO rules until the incoming order's quantity is fulfilled or all LMM orders have received their
allocated quantity.
Any remaining quantity is allocated pro-rata with configurable minimum allocation.
The Threshold Pro-Rata with LMM algorithm follows these stages to match trades:
	- Assigns a top order percent allocation (with min and/or max)
	- LMM set to configurable percent
	- Pro-rata with configurable minimum allocation
	- FIFO for any residual quantity
*/

type ThresholdProRataLMMEngine struct{
	OrderEngine
}

func (e *ThresholdProRataLMMEngine) run(symbol string) {

}



/*
AllocationEngine: Allocation Algorithm
The Allocation algorithm is an enhanced pro-rata algorithm that incorporates a priority (top order) to the first
incoming order that betters the market. If priority is established, the aggressor orders are first allocated to the
top order until the order's quantity is exhausted.
Additionally, there is a minimum pro-rata allocation parameter of two lots. All fills are rounded down to the
nearest integer and if an allocated trade quantity is less than two lots, it is rounded down to zero. Excess
lots are allocated FIFO.
The Allocation algorithm follows these stages to match trades:
	- Assigns a top order percent allocation
	- Pro-rata with minimum allocation of two lots
	- FIFO for any residual quantity
*/

type AllocationEngine struct{
	OrderEngine
}

func (e *AllocationEngine) run(symbol string) {

}



