package app

import (
	"net/http"
)

// Private Handlers
func (a *App) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(NewOrderParser)
	parser = a.Parse(parser,w,r)
	response, err:= a.Engine.NewOrder(account, parser)
	a.Respond(w,response,err)
	a.spot(r)
}

// Result of /order/status for the cancelled order.
func (a *App) CancelOrderHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(CancelOrderParser)
	parser = a.Parse(parser,w,r)
	response, err:= a.Engine.CancelOrder(account, parser)
	a.Respond(w,response,err)
	a.spot(r)
}

// Result of /order/status for the cancelled order.
func (a *App) CancelMultipleOrdersHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(CancelMultipleOrdersParser)
	parser = a.Parse(parser,w,r)
	response, err:= a.Engine.CancelMultipleOrders(account, parser)
	a.Respond(w,response,err)
	a.spot(r)
}

// Result of /order/status for the cancelled order.
func (a *App) CancelAllOrdersHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	response, err:= a.Engine.CancelAllOrders(account)
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) OrderHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(GetOrderParser)
	parser = a.Parse(parser,w,r)
	response, err := a.Engine.GetOrder(account, parser)
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) OrdersHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	response, err := a.Engine.GetOrders(account)
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) CreateOfferHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(NewOrderParser)
	parser = a.Parse(parser,w,r)
	response, err:= a.Engine.NewOffer(account, parser)
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) CancelOfferHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(CancelOrderParser)
	parser = a.Parse(parser,w,r)
	response, err:= a.Engine.CancelOffer(account, parser)
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) OfferHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(GetOrderParser)
	parser = a.Parse(parser,w,r)
	response, err := a.Engine.GetOffer(account, parser)
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) OffersHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	response, err := a.Engine.GetOffers(account)
	a.Respond(w,response,err)
	a.spot(r)
}


