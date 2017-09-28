package engine

import "github.com/thorad/simulacrum/app"

// Fundingbook
func (e *OrderEngine) GetFundingbook() app.GetFundingbookParser{

}

func (e *OrderEngine) GetLends() app.GetLendsParser{

}


// Orderbook
func (e *OrderEngine) GetOrderbook(parser app.GetOrderbookParser) app.GetOrderbookParser{

}

func (e *OrderEngine) GetOrderbooks()  app.GetOrderbooksParser{

}

// Trades
func (e *OrderEngine) GetTradeHistory(parser app.GetTradeHistoryParser)  app.GetOrderbookParser{

}

func (e *OrderEngine) GetTradeHistories() app.GetTradeHistoriesParser{

}


// Stats
func (e *OrderEngine) GetStat(parser app.GetStatParser) app.GetStatParser{

}

func (e *OrderEngine) GetStats() app.GetStatsParser{

}

// Tickers
func (e *OrderEngine) GetTicker(parser app.GetTickerParser) app.GetTickerParser{

}

func (e *OrderEngine) GetTickers() app.GetTickersParser{

}


// Assets & Pairs
func (e *OrderEngine) GetAssets() app.GetAssetsParser{

}

func (e *OrderEngine) GetAsset(parser app.GetAssetParser) app.GetAssetParser{

}

func (e *OrderEngine) GetPairs(parser app.GetPairParser) app.GetPairParser{

}

func (e *OrderEngine) GetDetailedPairs() app.GetDetailedPairsParser{

}

func (e *OrderEngine) GetPair(parser app.GetPairParser) app.GetPairParser{

}

func (e *OrderEngine) GetDetailedPair(parser app.GetDetailedPairParser) app.GetDetailedPairParser{

}