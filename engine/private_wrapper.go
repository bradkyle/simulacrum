package engine

import (
	"github.com/thorad/simulacrum/app"
	acc "github.com/thorad/simulacrum/account"
)

// Orders
func (e *OrderEngine) NewOrder(account acc.Account, parser app.NewOrderParser) app.NewOrderParser{

}

// return Result of /order/status for the cancelled order.
func (e *OrderEngine) CancelOrder(account acc.Account, parser app.CancelOrderParser) app.CancelOrderParser{

}

// return {"result":"Orders cancelled"}
func (e *OrderEngine) CancelMultipleOrders(account acc.Account, parser app.CancelOrderParser) app.CancelOrderParser{

}

// return {"result":"All orders cancelled"}
func (e *OrderEngine) CancelAllOrders(account acc.Account) app.CancelAllOrdersParser{

}

func (e *OrderEngine) GetOrder(account acc.Account, parser app.GetOrderParser) app.GetOrderParser{

}

func (e *OrderEngine) GetOrders(account acc.Account) app.GetOrdersParser{

}


//Offers
func (e *Engine) NewOffer(account acc.Account, parser app.NewOfferParser) app.NewOfferParser{

}

func (e *Engine) CancelOffer(account acc.Account, parser app.CancelOfferParser) app.CancelOfferParser{

}

func (e *Engine) GetOffer(account acc.Account, parser app.GetOfferParser) app.GetOfferParser{

}

func (e *Engine) GetOffers(account acc.Account) app.GetOffersParser{

}

func (e *Engine) CancelAllOffers(account acc.Account) app.CancelAllOffersParser{

}



