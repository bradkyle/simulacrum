package app

import (
	"net/http"
)

// Private Handlers
func (a *App) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	account := a.auth(r)
	parser := new(NewOrderParser)
	parser.Parse(w,r)
	response, err:= a.Engine.NewOrder(account, parser)
	write(w,response,err)
}

func (a *App) CancelOrderHandler(w http.ResponseWriter, r *http.Request) {
	account := a.auth(r)
	parser := new(CancelOrderParser)
	parser.Parse(w,r)
	response, err:= a.Engine.CancelOrder(account, parser)
	write(w,response,err)
}

func (a *App) OrderHandler(w http.ResponseWriter, r *http.Request) {
	account := a.auth(r)
	parser := new(GetOrderParser)
	parser.Parse(w,r)
	response, err := a.Engine.GetOrder(account, parser)
	write(w,response,err)
}

func (a *App) OrdersHandler(w http.ResponseWriter, r *http.Request) {
	account := a.auth(r)
	response, err := a.Engine.GetOrders(account)
	write(w,response,err)
}

func (a *App) CreateOfferHandler(w http.ResponseWriter, r *http.Request) {
	account := a.auth(r)
	parser := new(NewOrderParser)
	parser.Parse(w,r)
	response, err:= a.Engine.NewOffer(account, parser)
	write(w,response,err)
}

func (a *App) CancelOfferHandler(w http.ResponseWriter, r *http.Request) {
	account := a.auth(r)
	parser := new(CancelOrderParser)
	parser.Parse(w,r)
	response, err:= a.Engine.CancelOffer(account, parser)
	write(w,response,err)
}

func (a *App) OfferHandler(w http.ResponseWriter, r *http.Request) {
	account := a.auth(r)
	parser := new(GetOrderParser)
	parser.Parse(w,r)
	response, err := a.Engine.GetOffer(account, parser)
	write(w,response,err)
}

func (a *App) OffersHandler(w http.ResponseWriter, r *http.Request) {
	account := a.auth(r)
	response, err := a.Engine.GetOffers(account)
	write(w,response,err)
}


