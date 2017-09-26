package parse

type BitfinexOrderbook struct {
	Bids []struct {
		Price     string `json:"price"`
		Amount    string `json:"amount"`
		Timestamp string `json:"timestamp"`
	} `json:"bids"`
	Asks []struct {
		Price     string `json:"price"`
		Amount    string `json:"amount"`
		Timestamp string `json:"timestamp"`
	} `json:"asks"`
}

type GeminiOrderbook struct {
	Bids []struct {
		Price     string `json:"price"`
		Amount    string `json:"amount"`
		Timestamp string `json:"timestamp"`
	} `json:"bids"`
	Asks []struct {
		Price     string `json:"price"`
		Amount    string `json:"amount"`
		Timestamp string `json:"timestamp"`
	} `json:"asks"`
}

type KrakenOrderbook struct {
	Error  []interface{} `json:"error"`
	Result struct {
		BCHEUR struct {
			Asks [][]interface{} `json:"asks"`
			Bids [][]interface{} `json:"bids"`
		} `json:"BCHEUR"`
	} `json:"result"`
}

type PoloniexOrderbook struct {
	Asks     [][]int `json:"asks"`
	Bids     [][]int `json:"bids"`
	IsFrozen int     `json:"isFrozen"`
	Seq      int     `json:"seq"`
}


type PoloniexAllMarketsOrderbook struct {
	BTCAMP struct {
		Asks     [][]interface{} `json:"asks"`
		Bids     [][]interface{} `json:"bids"`
		IsFrozen string          `json:"isFrozen"`
		Seq      int             `json:"seq"`
	} `json:"BTC_AMP"`
}

type BitXSaOrderbook struct {
	Timestamp int64 `json:"timestamp"`
	Bids      []struct {
		Volume string `json:"volume"`
		Price  string `json:"price"`
	} `json:"bids"`
	Asks []struct {
		Volume string `json:"volume"`
		Price  string `json:"price"`
	} `json:"asks"`
}

type BtccOrderbook struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
	Date int64       `json:"date"`
}