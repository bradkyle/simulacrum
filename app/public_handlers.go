package app

import (
	"net/http"
)


// Public Handlers

// New Primary Entity Handlers
func (a *App) NewAccountHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	parser := new(NewAccountParser)
	parser = a.Parse(parser,w,r)
	response, err:= a.NewAccount(parser)
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) FundingbookHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response, err:= a.Engine.GetFundingbook()
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) LendsHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response, err:= a.Engine.GetLends()
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) OrderbookHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	parser := new(GetOrderbookParser)
	parser = a.Parse(parser,w,r)
	response, err:= a.Engine.GetOrderbook(parser)
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) OrderbooksHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response, err:= a.Engine.GetOrderbooks()
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) TradesHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response, err:= a.Engine.GetTrades()
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) StatHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	parser := new(GetStatParser)
	parser = a.Parse(parser,w,r)
	response, err:= a.Engine.GetStat(parser)
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) StatsHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response, err:= a.Engine.GetStats()
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) TickerHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	parser := new(GetTickerParser)
	parser = a.Parse(parser,w,r)
	response, err:= a.Engine.GetTicker(parser)
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) TickersHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response, err:= a.Engine.GetTickers()
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) AssetsHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response, err:= a.Engine.GetAssets()
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) AssetHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	parser := new(GetAssetParser)
	parser = a.Parse(parser,w,r)
	response, err:= a.Engine.GetAsset(parser)
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) PairsHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response, err:= a.Engine.GetPairs()
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) DetailedPairsHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response, err:= a.Engine.GetDetailedPairs()
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) PairHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	parser := new(GetPairParser)
	parser = a.Parse(parser,w,r)
	response, err:= a.Engine.GetPair(parser)
	a.Respond(w,response,err)
	a.spot(r)
}

func (a *App) DetailedPairHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	parser := new(GetDetailedPairParser)
	parser = a.Parse(parser,w,r)
	response, err:= a.Engine.GetDetailedPair(parser)
	a.Respond(w,response,err)
	a.spot(r)
}

