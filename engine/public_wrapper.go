package engine

import "github.com/thorad/simulacrum/parse"

// Fundingbook
func (e *OrderEngine) GetFundingbook() parse.GetFundingbookParser{

}

func (e *OrderEngine) GetLends() parse.GetLendsParser{

}


// Orderbook
func (e *OrderEngine) GetOrderbook(parser parse.GetOrderbookParser) parse.GetOrderbookParser{

}

func (e *OrderEngine) GetOrderbooks()  parse.GetOrderbooksParser{

}

// Trades
func (e *OrderEngine) GetTradeHistory(parser parse.GetTradeHistoryParser)  parse.GetOrderbookParser{

}

func (e *OrderEngine) GetTradeHistories() parse.GetTradeHistoriesParser{

}


// Stats
func (e *OrderEngine) GetStat(parser parse.GetStatParser) parse.GetStatParser{

}

func (e *OrderEngine) GetStats() parse.GetStatsParser{

}

// Tickers
func (e *OrderEngine) GetTicker(parser parse.GetTickerParser) parse.GetTickerParser{

}

func (e *OrderEngine) GetTickers() parse.GetTickersParser{

}


// Assets & Pairs
func (e *OrderEngine) GetAssets() parse.GetAssetsParser{

}

func (e *OrderEngine) GetAsset(parser parse.GetAssetParser) parse.GetAssetParser{

}

func (e *OrderEngine) GetPairs(parser parse.GetPairParser) parse.GetPairParser{

}

func (e *OrderEngine) GetDetailedPairs() parse.GetDetailedPairsParser{

}

func (e *OrderEngine) GetPair(parser parse.GetPairParser) parse.GetPairParser{

}

func (e *OrderEngine) GetDetailedPair(parser parse.GetDetailedPairParser) parse.GetDetailedPairParser{

}