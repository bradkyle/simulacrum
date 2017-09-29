package engine

import (
	"gopkg.in/redsync.v1"
	"time"
	"sync"
	"github.com/garyburd/redigo/redis"
	"github.com/thorad/simulacrum/account"
)

type EngineCore interface{
	Switch()
}

type Engine struct{
	lockMap       map[string]*Locker
	redsync       *redsync.Redsync
}



type MatchingEngineType int
const(
	PriceTimePriority MatchingEngineType = iota
	ProRata
)


type OrderEngine struct{
	Engine
	LastPrice     float64
	queue		*OrderQueue
	store           *OrderHash
}


type OrderEngineCore interface{

}

func (e *OrderEngine) publishOrderbook(symbol string, buyTop *Node, sellTop *Node) {
	bid := &Order{Symbol: symbol}
	ask := &Order{Symbol: symbol}

	if buyTop != nil {
		bid = e.store.get(buyTop.Lookup)
	}
	if sellTop != nil {
		ask = e.store.get(sellTop.Lookup)
	}
	go e.Publish(bid, ask)
}

func (e *OrderEngine) Publish(bid *Order, ask *Order){

}

func (e *OrderEngine) Switch(){

}






type FundingEngine struct{

}

type FundingEngineCore interface{

}

func (e *FundingEngine) publishOfferbook() {

}

func (e *FundingEngine) Publish(){

}

func (e *FundingEngine) Switch(){

}





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

