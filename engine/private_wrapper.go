package engine

import (
	"github.com/thorad/simulacrum/app"
	acc "github.com/thorad/simulacrum/account"
)

// Orders
func (e *Engine) NewOrder(account acc.Account, parser app.NewOrderParser) (){
	locker := o.getLocker(payload.Ticker)
	err := locker.Lock()
	if err != nil {
		return err
	}
	defer locker.Unlock()

	o.orderStore.set(order.lookup(), order)
	o.orderQueue.Enqueue(order.Intent+order.Ticker, &heap.Node{
		Value:  order.price(),
		Lookup: order.lookup(),
	})

	o.run(payload.Ticker)
	return nil
}

func (e *Engine) CancelOrder(account acc.Account, parser app.CancelOrderParser) (){

}

func (e *Engine) GetOrder(account acc.Account, parser app.GetOrderParser) (){

}

func (e *Engine) GetOrders(account acc.Account) (){

}


//Offers
func (e *Engine) NewOffer(account acc.Account, parser app.NewOfferParser) (){
	locker := o.getLocker(payload.Ticker)
	err := locker.Lock()
	if err != nil {
		return err
	}
	defer locker.Unlock()

	o.orderStore.set(order.lookup(), order)
	o.orderQueue.Enqueue(order.Intent+order.Ticker, &heap.Node{
		Value:  order.price(),
		Lookup: order.lookup(),
	})

	o.run(payload.Ticker)
	return nil
}

func (e *Engine) CancelOffer(account acc.Account, parser app.CancelOfferParser) (){

}

func (e *Engine) GetOffer(account acc.Account, parser app.GetOfferParser) (){

}

func (e *Engine) GetOffers(account acc.Account) (){

}


