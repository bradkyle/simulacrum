package parse


type GeminiTicker struct {
	Ask    string `json:"ask"`
	Bid    string `json:"bid"`
	Last   string `json:"last"`
	Volume struct {
		BTC       string `json:"BTC"`
		USD       string `json:"USD"`
		Timestamp int64  `json:"timestamp"`
	} `json:"volume"`
}

type PoloniexTicker struct {
	BTCLTC struct {
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
	} `json:"BTC_LTC"`
}

type KrakenTicker struct {
	Error  []interface{} `json:"error"`
	Result struct {
		BCHEUR struct {
			A []string `json:"a"`
			B []string `json:"b"`
			C []string `json:"c"`
			V []string `json:"v"`
			P []string `json:"p"`
			T []int    `json:"t"`
			L []string `json:"l"`
			H []string `json:"h"`
			O string   `json:"o"`
		} `json:"BCHEUR"`
	} `json:"result"`
}

type BitfinexTicker struct {
	Mid       string `json:"mid"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	LastPrice string `json:"last_price"`
	Low       string `json:"low"`
	High      string `json:"high"`
	Volume    string `json:"volume"`
	Timestamp string `json:"timestamp"`
}

type HitBTC struct {
	Last        string `json:"last"`
	Bid         string `json:"bid"`
	Ask         string `json:"ask"`
	High        string `json:"high"`
	Low         string `json:"low"`
	Volume      string `json:"volume"`
	Open        string `json:"open"`
	VolumeQuote string `json:"volume_quote"`
	Timestamp   int64  `json:"timestamp"`
}

type BithumbTicker struct {
	Status string `json:"status"`
	Data   struct {
		OpeningPrice string `json:"opening_price"`
		ClosingPrice string `json:"closing_price"`
		MinPrice     string `json:"min_price"`
		MaxPrice     string `json:"max_price"`
		AveragePrice string `json:"average_price"`
		UnitsTraded  string `json:"units_traded"`
		Volume1Day   string `json:"volume_1day"`
		Volume7Day   string `json:"volume_7day"`
		Date         int64  `json:"date"`
	} `json:"data"`
}

type CoinoneTicker struct {
	Volume    string `json:"volume"`
	Last      string `json:"last"`
	Timestamp string `json:"timestamp"`
	High      string `json:"high"`
	Result    string `json:"result"`
	ErrorCode string `json:"errorCode"`
	First     string `json:"first"`
	Low       string `json:"low"`
	Currency  string `json:"currency"`
}

type BitstampTicker struct {
	High      string  `json:"high"`
	Last      string  `json:"last"`
	Timestamp string  `json:"timestamp"`
	Bid       string  `json:"bid"`
	Vwap      string  `json:"vwap"`
	Volume    string  `json:"volume"`
	Low       string  `json:"low"`
	Ask       string  `json:"ask"`
	Open      float64 `json:"open"`
}

type BitFlyerApi struct {
	ProductCode     string  `json:"product_code"`
	Timestamp       string  `json:"timestamp"`
	TickID          int     `json:"tick_id"`
	BestBid         int     `json:"best_bid"`
	BestAsk         int     `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     int     `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   int     `json:"total_ask_depth"`
	Ltp             int     `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

type OkCoinTicker struct {
	Date   string `json:"date"`
	Ticker struct {
		Buy  string `json:"buy"`
		High string `json:"high"`
		Last string `json:"last"`
		Low  string `json:"low"`
		Sell string `json:"sell"`
		Vol  string `json:"vol"`
	} `json:"ticker"`
}

type KorbitTicker struct {
	Timestamp int64  `json:"timestamp"`
	Last      string `json:"last"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	Low       string `json:"low"`
	High      string `json:"high"`
	Volume    string `json:"volume"`
}

type LiquiTicker struct {
	EthBtc struct {
		High    float64 `json:"high"`
		Low     float64 `json:"low"`
		Avg     float64 `json:"avg"`
		Vol     float64 `json:"vol"`
		VolCur  float64 `json:"vol_cur"`
		Last    float64 `json:"last"`
		Buy     float64 `json:"buy"`
		Sell    float64 `json:"sell"`
		Updated int     `json:"updated"`
	} `json:"eth_btc"`
}

type GatecoinTicker struct {
	Tickers []struct {
		CurrencyPair   string  `json:"currencyPair"`
		Open           int     `json:"open"`
		Last           int     `json:"last"`
		LastQ          float64 `json:"lastQ"`
		High           int     `json:"high"`
		Low            int     `json:"low"`
		Volume         int     `json:"volume"`
		Bid            int     `json:"bid"`
		BidQ           float64 `json:"bidQ"`
		Ask            int     `json:"ask"`
		AskQ           int     `json:"askQ"`
		Vwap           int     `json:"vwap"`
		CreateDateTime string  `json:"createDateTime"`
	} `json:"tickers"`
}

type Btc38Ticker struct {
	Ticker struct {
		High int     `json:"high"`
		Low  int     `json:"low"`
		Last int     `json:"last"`
		Vol  float64 `json:"vol"`
		Buy  int     `json:"buy"`
		Sell int     `json:"sell"`
	} `json:"ticker"`
}

