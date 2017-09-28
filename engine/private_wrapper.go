package engine

import (
	"github.com/thorad/simulacrum/app"
	acc "github.com/thorad/simulacrum/account"
)

// Orders
// return order_id [int64] An order object containing the order’s ID as well as all the information provided by /order/status
func (e *OrderEngine) NewOrder(account acc.Account, parser app.NewOrderParser) (){
	locker := e.getLocker(parser.Symbol)
	err := locker.Lock()
	if err != nil {
		return err
	}
	defer locker.Unlock()

	// todo convert parser into order
	order := Order{}

	e.store.set(order.index(), order)
	e.queue.Enqueue(order.Side()+order.Symbol, &Node{
		Value:  order.price(),
		Lookup: order.index(),
	})

	e.run(parser.Symbol)
	return nil
}

// return Result of /order/status for the cancelled order.
func (e *OrderEngine) CancelOrder(account acc.Account, parser app.CancelOrderParser) (){

}

// return {"result":"Orders cancelled"}
func (e *OrderEngine) CancelMultipleOrders(account acc.Account, parser app.CancelOrderParser) (){

}

// return {"result":"All orders cancelled"}
func (e *OrderEngine) CancelAllOrders(account acc.Account, parser app.CancelOrderParser) (){

}

func (e *OrderEngine) GetOrder(account acc.Account, parser app.GetOrderParser) (){

}

func (e *OrderEngine) GetOrders(account acc.Account) (){

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


