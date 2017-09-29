package engine

import (
	"gopkg.in/redsync.v1"
	"time"
	"sync"
	"github.com/garyburd/redigo/redis"
	"github.com/thorad/simulacrum/account"
	"fmt"
)

type EngineCore interface{
	Switch()
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


type OrderEngineCore interface{

}

func (e *OrderEngine) Publish(symbol string, bid *Order, ask *Order){

}

func (e *OrderEngine) Switch(){


}

func (e *OrderEngine) negotiate(buy *Order, sell *Order) float64{
	bKind := buy.Kind
	sKind := sell.Kind

	if bKind == Market && sKind == Limit{
		e.LastPrice = sell.price()

	} else if sKind == Market && bKind == Limit {
		e.LastPrice = buy.price()

	} else if bKind == Limit && sKind == Limit {
		e.LastPrice = sell.price()

	} // else if both market, use last price

	return e.LastPrice
}


func (e *OrderEngine) trade(buy *Trade, sell *Trade){

}



func NewEngine(pool *redis.Pool) *Engine {
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

