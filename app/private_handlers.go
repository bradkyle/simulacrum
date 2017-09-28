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
	response := a.orderEngine.NewOrder(account, parser)
	a.Respond(w,response)
	a.spot(r)
}

// Result of /order/status for the cancelled order.
func (a *App) CancelOrderHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(CancelOrderParser)
	parser = a.Parse(parser,w,r)
	response := a.orderEngine.CancelOrder(account, parser)
	a.Respond(w,response)
	a.spot(r)
}

// Result of /order/status for the cancelled order.
func (a *App) CancelMultipleOrdersHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(CancelMultipleOrdersParser)
	parser = a.Parse(parser,w,r)
	response := a.orderEngine.CancelMultipleOrders(account, parser)
	a.Respond(w,response)
	a.spot(r)
}

// Result of /order/status for the cancelled order.
func (a *App) CancelAllOrdersHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	response := a.orderEngine.CancelAllOrders(account)
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) OrderHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(GetOrderParser)
	parser = a.Parse(parser,w,r)
	response := a.orderEngine.GetOrder(account, parser)
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) OrdersHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	response := a.orderEngine.GetOrders(account)
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) CreateOfferHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(NewOrderParser)
	parser = a.Parse(parser,w,r)
	response := a.orderEngine.NewOffer(account, parser)
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) CancelOfferHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(CancelOrderParser)
	parser = a.Parse(parser,w,r)
	response := a.orderEngine.CancelOffer(account, parser)
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) OfferHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	parser := new(GetOrderParser)
	parser = a.Parse(parser,w,r)
	response := a.orderEngine.GetOffer(account, parser)
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) OffersHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	account := a.auth(r)
	response := a.orderEngine.GetOffers(account)
	a.Respond(w,response)
	a.spot(r)
}


