package engine

import (
	"gopkg.in/redsync.v1"
	"time"
	"sync"
	"github.com/garyburd/redigo/redis"
	"github.com/thorad/simulacrum/account"
)

type EngineCore interface{

}

type Engine struct{
	lockMap       map[string]*Locker
	redsync       *redsync.Redsync
}

type OrderEngine struct{
	Engine
	LastPrice     float64
	queue		*OrderQueue
	store           *OrderHash
}

type FundingEngine struct{

}


type MatchingEngineType int
const(
	PriceTimePriority MatchingEngineType = iota
)


func NewEngine(pool *redis.Pool, kind MatchingEngineType) *Engine {
	return &Engine{
		lockMap:      	make(map[string]*Locker),
		redsync:       	redsync.New([]redsync.Pool{pool}),
	}
}


func (e *Engine) getLocker(symbol string) *Locker {
	if nil != e.lockMap[symbol] {
		return e.lockMap[symbol]
	} else {
		redLockMutex := e.redsync.NewMutex("orderbook_service" + symbol)
		redsync.SetRetryDelay(5 * time.Millisecond).Apply(redLockMutex)
		redsync.SetExpiry(500 * time.Millisecond).Apply(redLockMutex)
		redsync.SetTries(50).Apply(redLockMutex)

		e.lockMap[symbol] = &Locker{
			name:    "orderbook_service" + symbol,
			mutLock: &sync.Mutex{},
			redLock: redLockMutex,
		}
		return e.lockMap[symbol]
	}
}

func (e *OrderEngine) negotiate(b *Order, s *Order) float64 {
	bKind := b.Kind
	sKind := s.Kind

	if bKind == "MARKET" && sKind == "LIMIT" {
		e.LastPrice = s.price()

	} else if sKind == "MARKET" && bKind == "LIMIT" {
		e.LastPrice = b.price()

	} else if bKind == "LIMIT" && sKind == "LIMIT" {
		e.LastPrice = s.price()

	} // else if both market, use last price
	return e.LastPrice
}

// does not match orders as soon as they enter
func (e *OrderEngine) run(symbol string) {
	buyTop := e.queue.Peek("BUY" + symbol)
	sellTop := e.queue.Peek("SELL" + symbol)

	for buyTop != nil && sellTop != nil && buyTop.Value >= sellTop.Value {

		buy := e.store.get(buyTop.Lookup)
		sell := e.store.get(sellTop.Lookup)
		price := e.negotiate(buy, sell)

		if buy.Amount == sell.Amount {
			e.queue.Dequeue("BUY" + symbol)
			e.queue.Dequeue("SELL" + symbol)

			go account.ProcessTrade(buy.fill(price), sell.fill(price))

			e.store.remove(buyTop.Lookup)
			e.store.remove(sellTop.Lookup)

		} else if buy.Amount < sell.Amount {
			e.queue.Dequeue("BUY" + symbol)
			remainderSell := sell.Amount - buy.Amount

			go account.ProcessTrade(buy.fill(price), sell.partialFill(price, remainderSell))

			e.store.remove(buyTop.Lookup)

		} else if buy.Amount > sell.Amount {
			e.queue.Dequeue("SELL" + symbol)
			remainderBuy := buy.Amount - sell.Amount

			go account.ProcessTrade(sell.fill(price), sell.partialFill(price, remainderBuy))

			e.store.remove(sellTop.Lookup)
		}

		buyTop = e.queue.Peek("BUY" + symbol)
		sellTop = e.queue.Peek("SELL" + symbol)
	}

	e.publishOrderbook(symbol, buyTop, sellTop)
}

func (o *OrderEngine) publishOrderbook(symbol string, buyTop *Node, sellTop *Node) {
	bid := &Order{Symbol: symbol}
	ask := &Order{Symbol: symbol}

	if buyTop != nil {
		bid = o.store.get(buyTop.Lookup)
	}
	if sellTop != nil {
		ask = o.store.get(sellTop.Lookup)
	}
	go o.publishBidAsk(bid, ask)
}
