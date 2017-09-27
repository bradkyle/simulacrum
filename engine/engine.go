package engine

import (
	"gopkg.in/redsync.v1"
	"time"
	"sync"
)

type EngineCore interface{

}

type Engine struct{

}

type FundingEngine struct{

}

type Engine struct{

}

type MatchingEngineType int
const(
	PriceTimePriority MatchingEngineType = iota
)


func NewEngine(pool *redis.Pool) *Orderbook {
	return &Orderbook{
		lockMap:      	make(map[string]*Locker),
		redsync:       	redsync.New([]redsync.Pool{pool}),
		BuyQueue:    	NewOrderQueue(pool),
		SellQueue:   	NewOrderQueue(pool),
	}
}


func (e *Engine) getLocker(ticker string) *Locker {
	if nil != o.lockMap[ticker] {
		return o.lockMap[ticker]
	} else {
		redLockMutex := o.redsync.NewMutex("orderbook_service" + ticker)
		redsync.SetRetryDelay(5 * time.Millisecond).Apply(redLockMutex)
		redsync.SetExpiry(500 * time.Millisecond).Apply(redLockMutex)
		redsync.SetTries(50).Apply(redLockMutex)

		o.lockMap[ticker] = &Locker{
			name:    "orderbook_service" + ticker,
			mutLock: &sync.Mutex{},
			redLock: redLockMutex,
		}
		return o.lockMap[ticker]
	}
}

func (e *Engine) negotiate(b *Order, s *Order) float64 {
	bKind := b.Kind
	sKind := s.Kind

	if bKind == "MARKET" && sKind == "LIMIT" {
		o.lastPrice = s.price()

	} else if sKind == "MARKET" && bKind == "LIMIT" {
		o.lastPrice = b.price()

	} else if bKind == "LIMIT" && sKind == "LIMIT" {
		o.lastPrice = s.price()

	} // else if both market, use last price

	return o.lastPrice
}

// does not match orders as soon as they enter
func (o *Orderbook) run(ticker string) {
	buyTop := o.orderQueue.Peek("BUY" + ticker)
	sellTop := o.orderQueue.Peek("SELL" + ticker)

	for buyTop != nil && sellTop != nil && buyTop.Value >= sellTop.Value {

		buy := o.orderStore.get(buyTop.Lookup)
		sell := o.orderStore.get(sellTop.Lookup)
		price := o.negotiate(buy, sell)

		if buy.Amount == sell.Amount {
			o.orderQueue.Dequeue("BUY" + ticker)
			o.orderQueue.Dequeue("SELL" + ticker)

			go o.handleTrade(buy.fill(price), sell.fill(price))

			o.orderStore.remove(buyTop.Lookup)
			o.orderStore.remove(sellTop.Lookup)

		} else if buy.Amount < sell.Amount {
			o.orderQueue.Dequeue("BUY" + ticker)
			remainderSell := sell.Amount - buy.Amount

			go o.handleTrade(buy.fill(price), sell.partialFill(price, remainderSell))

			o.orderStore.remove(buyTop.Lookup)

		} else if buy.Amount > sell.Amount {
			o.orderQueue.Dequeue("SELL" + ticker)
			remainderBuy := buy.Amount - sell.Amount

			go o.handleTrade(sell.fill(price), sell.partialFill(price, remainderBuy))

			o.orderStore.remove(sellTop.Lookup)
		}

		buyTop = o.orderQueue.Peek("BUY" + ticker)
		sellTop = o.orderQueue.Peek("SELL" + ticker)
	}

	o.publishOrderbook(ticker, buyTop, sellTop)
}

func (o *Orderbook) publishOrderbook(ticker string, buyTop *heap.Node, sellTop *heap.Node) {
	bid := &Order{Ticker: ticker}
	ask := &Order{Ticker: ticker}

	if buyTop != nil {
		bid = o.orderStore.get(buyTop.Lookup)
	}
	if sellTop != nil {
		ask = o.orderStore.get(sellTop.Lookup)
	}
	go o.publishBidAsk(bid, ask)
}
