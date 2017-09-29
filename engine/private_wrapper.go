package engine

import (
	"github.com/thorad/simulacrum/parse"
	acc "github.com/thorad/simulacrum/account"
)

// Orders
func (e *OrderEngine) NewOrder(account acc.Account, parser parse.NewOrderParser) parse.NewOrderParser{

}

// return Result of /order/status for the cancelled order.
func (e *OrderEngine) CancelOrder(account acc.Account, parser parse.CancelOrderParser) parse.CancelOrderParser{

}

// return {"result":"Orders cancelled"}
func (e *OrderEngine) CancelMultipleOrders(account acc.Account, parser parse.CancelOrderParser) parse.CancelOrderParser{

}

// return {"result":"All orders cancelled"}
func (e *OrderEngine) CancelAllOrders(account acc.Account) parse.CancelAllOrdersParser{

}

func (e *OrderEngine) GetOrder(account acc.Account, parser parse.GetOrderParser) parse.GetOrderParser{

}

func (e *OrderEngine) GetOrders(account acc.Account) parse.GetOrdersParser{

}


//Offers
func (e *Engine) NewOffer(account acc.Account, parser parse.NewOfferParser) parse.NewOfferParser{

}

func (e *Engine) CancelOffer(account acc.Account, parser parse.CancelOfferParser) parse.CancelOfferParser{

}

func (e *Engine) GetOffer(account acc.Account, parser parse.GetOfferParser) parse.GetOfferParser{

}

func (e *Engine) GetOffers(account acc.Account) parse.GetOffersParser{

}

func (e *Engine) CancelAllOffers(account acc.Account) parse.CancelAllOffersParser{

}



