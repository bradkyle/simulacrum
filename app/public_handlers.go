package app

import (
	"net/http"
)


// Public Handlers

// New Primary Entity Handlers
func (a *App) NewAccountHandler(w http.ResponseWriter, r *http.Request) {
	parser := new(NewAccountParser)
	parser.Parse(w,r)
	response, err:= a.NewAccount(parser)
	write(w,response,err)
	a.spot(r)
}

func (a *App) FundingbookHandler(w http.ResponseWriter, r *http.Request) {
	response, err:= a.Engine.GetFundingbook()
	write(w,response,err)
	a.spot(r)
}

func (a *App) LendsHandler(w http.ResponseWriter, r *http.Request) {
	response, err:= a.Engine.GetLends()
	write(w,response,err)
	a.spot(r)
}

func (a *App) OrderbookHandler(w http.ResponseWriter, r *http.Request) {
	parser := new(GetOrderbookParser)
	parser.Parse(w,r)
	response, err:= a.Engine.GetOrderbook(parser)
	write(w,response,err)
	a.spot(r)
}

func (a *App) OrderbooksHandler(w http.ResponseWriter, r *http.Request) {
	response, err:= a.Engine.GetOrderbooks()
	write(w,response,err)
	a.spot(r)
}

func (a *App) TradesHandler(w http.ResponseWriter, r *http.Request) {
	response, err:= a.Engine.GetTrades()
	write(w,response,err)
	a.spot(r)
}

func (a *App) StatHandler(w http.ResponseWriter, r *http.Request) {
	parser := new(GetStatParser)
	parser.Parse(w,r)
	response, err:= a.Engine.GetStat(parser)
	write(w,response,err)
	a.spot(r)
}

func (a *App) StatsHandler(w http.ResponseWriter, r *http.Request) {
	response, err:= a.Engine.GetStats()
	write(w,response,err)
	a.spot(r)
}

func (a *App) TickerHandler(w http.ResponseWriter, r *http.Request) {
	parser := new(GetTickerParser)
	parser.Parse(w,r)
	response, err:= a.Engine.GetTicker(parser)
	write(w,response,err)
	a.spot(r)
}

func (a *App) TickersHandler(w http.ResponseWriter, r *http.Request) {
	response, err:= a.Engine.GetTickers()
	write(w,response,err)
	a.spot(r)
}

func (a *App) AssetsHandler(w http.ResponseWriter, r *http.Request) {
	response, err:= a.Engine.GetAssets()
	write(w,response,err)
	a.spot(r)
}

func (a *App) AssetHandler(w http.ResponseWriter, r *http.Request) {
	parser := new(GetAssetParser)
	parser.Parse(w,r)
	response, err:= a.Engine.GetAsset(parser)
	write(w,response,err)
	a.spot(r)
}

func (a *App) PairsHandler(w http.ResponseWriter, r *http.Request) {
	response, err:= a.Engine.GetPairs()
	write(w,response,err)
	a.spot(r)
}

func (a *App) DetailedPairsHandler(w http.ResponseWriter, r *http.Request) {
	response, err:= a.Engine.GetDetailedPairs()
	write(w,response,err)
	a.spot(r)
}

func (a *App) PairHandler(w http.ResponseWriter, r *http.Request) {
	parser := new(GetPairParser)
	parser.Parse(w,r)
	response, err:= a.Engine.GetPair(parser)
	write(w,response,err)
	a.spot(r)
}

func (a *App) DetailedPairHandler(w http.ResponseWriter, r *http.Request) {
	parser := new(GetDetailedPairParser)
	parser.Parse(w,r)
	response, err:= a.Engine.GetDetailedPair(parser)
	write(w,response,err)
	a.spot(r)
}

