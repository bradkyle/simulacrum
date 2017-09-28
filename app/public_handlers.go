package app

import (
	"net/http"
	acc "github.com/thorad/simulacrum/account"
)


// Public Handlers

// New Primary Entity Handlers
func (a *App) NewAccountHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	parser := new(NewAccountParser)
	parser = a.Parse(parser,w,r)
	response := acc.NewAccount(parser)
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) FundingbookHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response := a.orderEngine.GetFundingbook()
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) LendsHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response := a.orderEngine.GetLends()
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) OrderbookHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	parser := new(GetOrderbookParser)
	parser = a.Parse(parser,w,r)
	response := a.orderEngine.GetOrderbook(parser)
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) OrderbooksHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response := a.orderEngine.GetOrderbooks()
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) TradeHistoriesHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response := a.orderEngine.GetTradeHistories()
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) TradeHistoryHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response := a.orderEngine.GetTradeHistory()
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) StatHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	parser := new(GetStatParser)
	parser = a.Parse(parser,w,r)
	response := a.orderEngine.GetStat(parser)
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) StatsHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response := a.orderEngine.GetStats()
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) TickerHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	parser := new(GetTickerParser)
	parser = a.Parse(parser,w,r)
	response := a.orderEngine.GetTicker(parser)
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) TickersHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response := a.orderEngine.GetTickers()
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) AssetsHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response := a.orderEngine.GetAssets()
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) AssetHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	parser := new(GetAssetParser)
	parser = a.Parse(parser,w,r)
	response := a.orderEngine.GetAsset(parser)
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) PairsHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response := a.orderEngine.GetPairs()
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) DetailedPairsHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	response := a.orderEngine.GetDetailedPairs()
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) PairHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	parser := new(GetPairParser)
	parser = a.Parse(parser,w,r)
	response := a.orderEngine.GetPair(parser)
	a.Respond(w,response)
	a.spot(r)
}

func (a *App) DetailedPairHandler(w http.ResponseWriter, r *http.Request) {
	a.rule(w,r)
	parser := new(GetDetailedPairParser)
	parser = a.Parse(parser,w,r)
	response := a.orderEngine.GetDetailedPair(parser)
	a.Respond(w,response)
	a.spot(r)
}

