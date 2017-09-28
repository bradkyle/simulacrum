package engine

import "github.com/thorad/simulacrum/app"

// Fundingbook
func (e *Engine) GetFundingbook() (){

}

func (e *Engine) GetLends() (){

}


// Orderbook
func (e *Engine) GetOrderbook(parser app.GetOrderbookParser) (){

}

func (e *Engine) GetOrderbooks() (){

}

// Trades
func (e *Engine) GetTradeHistory(parser app.GetTradeHistoryParser) (){

}

func (e *Engine) GetTradeHistories(parser app.GetTradeHistoryParser) (){

}


// Stats
func (e *Engine) GetStat(parser app.GetStatParser) (){

}

func (e *Engine) GetStats() (){

}

// Tickers
func (e *Engine) GetTicker(parser app.GetTickerParser) (){

}

func (e *Engine) GetTickers() (){

}


// Assets & Pairs
func (e *Engine) GetAssets() (){

}

func (e *Engine) GetAsset(parser app.GetAssetParser) (){

}

func (e *Engine) GetPairs() (){

}

func (e *Engine) GetDetailedPairs() (){

}

func (e *Engine) GetPair(parser app.GetPairParser) (){

}

func (e *Engine) GetDetailedPair(parser app.GetDetailedPairParser) (){

}