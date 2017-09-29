package engine

import (
	"github.com/thorad/simulacrum/account"
	"time"
)

// Continuous Trading	----------------------------------------------------------------------------------------------->

/*
FifoTwoTradeEngine: Price/Time Algorithm
The FIFO algorithm uses price and time as the only criteria for filling an order. In this algorithm,
all orders at the same price level are filled according to time priority; the first order at a price
level is the first order matched.
It is important to note that an order loses order priority and is re-queued when changed in any of the following ways:
 - Increase the quantity
 - Change the price
 - Change the account number
In FifoTwoTradeEngine the subsequent trades are split into two seperate trades.
*/
type FifoTwoTradeEngine struct{
	OrderEngine
}

func (e *FifoTwoTradeEngine) run(symbol string) {

}



/*
FifoTwoTradeLMMEngine: Price/Time Algorithm
The FIFO with LMM algorithm is an enhanced FIFO algorithm that allows for LMM allocations prior to the
FIFO allocations. LMMs are each allocated a configurable percentage of an aggressor order before the
remaining quantity is allocated FIFO.
The FIFO with LMM algorithm follows these stages to match trades:
	- LMM set to configurable percent
	- FIFO allocation
*/

type FifoTwoTradeLMMEngine struct{
	OrderEngine
}

func (e *FifoTwoTradeLMMEngine) run(symbol string) {

}



/*
FifoTwoTradeTopLMMEngine: Price/Time Algorithm
The FIFO with LMM algorithm is an enhanced FIFO algorithm that allows for LMM allocations prior
to the FIFO allocations. Additionally, this algorithm incorporates a priority (top order) to the
first incoming order that betters the market. If priority is established, the aggressor orders are
first allocated to the top order until the order's quantity is exhausted.
The FIFO with LMM algorithm follows these stages to match trades:
	- Assigns a top order percent allocation that betters the market
	- LMM set to configurable percent
	- FIFO allocation
*/

type FifoTwoTradeTopLMMEngine struct{
	OrderEngine
}

func (e *FifoTwoTradeTopLMMEngine) run(symbol string) {

}



/*
FifoOneTradeEngine: Price/Time Algorithm
The FIFO algorithm uses price and time as the only criteria for filling an order. In this algorithm,
all orders at the same price level are filled according to time priority; the first order at a price
level is the first order matched.
It is important to note that an order loses order priority and is re-queued when changed in any of the following ways:
	- Increase the quantity
	- Change the price
	- Change the account number
In FifoOneTradeEngine only one trade is made accounting for both the buy and sell order.
*/

type FifoOneTradeEngine struct{
	OrderEngine
}

func (e *FifoOneTradeEngine) run(symbol string) {
	
}


/*
FifoOneTradeLMMEngine: Price/Time Algorithm
The FIFO with LMM algorithm is an enhanced FIFO algorithm that allows for LMM allocations prior to the
FIFO allocations. LMMs are each allocated a configurable percentage of an aggressor order before the
remaining quantity is allocated FIFO.
The FIFO with LMM algorithm follows these stages to match trades:
	- LMM set to configurable percent
	- FIFO allocation
*/

type FifoOneTradeLMMEngine struct{
	OrderEngine
}

func (e *FifoOneTradeLMMEngine) run(symbol string) {

}



/*
FifoOneTradeTopLMMEngine: Price/Time Algorithm
The FIFO with LMM algorithm is an enhanced FIFO algorithm that allows for LMM allocations prior
to the FIFO allocations. Additionally, this algorithm incorporates a priority (top order) to the
first incoming order that betters the market. If priority is established, the aggressor orders are
first allocated to the top order until the order's quantity is exhausted.
The FIFO with LMM algorithm follows these stages to match trades:
	- Assigns a top order percent allocation that betters the market
	- LMM set to configurable percent
	- FIFO allocation
*/

type FifoOneTradeTopLMMEngine struct{
	OrderEngine
}

func (e *FifoOneTradeTopLMMEngine) run(symbol string) {

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



// Auction Trading ---------------------------------------------------------------------------------------------------->







// Broker Based ------------------------------------------------------------------------------------------------------->







// Peer to Peer ------------------------------------------------------------------------------------------------------->