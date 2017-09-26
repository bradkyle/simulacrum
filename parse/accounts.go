package parse


type BitfinexAccountInfo []struct {
	MakerFees string `json:"maker_fees"`
	TakerFees string `json:"taker_fees"`
	Fees      []struct {
		Pairs     string `json:"pairs"`
		MakerFees string `json:"maker_fees"`
		TakerFees string `json:"taker_fees"`
	} `json:"fees"`
}

type BitfinexWalletBalances []struct {
	Type      string `json:"type"`
	Currency  string `json:"currency"`
	Amount    string `json:"amount"`
	Available string `json:"available"`
}

type GeminiBalance []struct {
	Type                   string `json:"type"`
	Currency               string `json:"currency"`
	Amount                 string `json:"amount"`
	Available              string `json:"available"`
	AvailableForWithdrawal string `json:"availableForWithdrawal"`
}

type KrakenBalances struct {
	Error  []interface{} `json:"error"`
	Result struct {
		       A string `json:"eb"` // equivalent balance (combined balance of all currencies)
		       B string `json:"tb"` // trade balance (combined balance of all equity currencies)
		       C string `json:"m"`  // margin amount of open positions
		       V string `json:"n"`  // volume array
		       P string `json:"c"`  // unrealized net profit/loss of open positions
		       T int    `json:"v"`  // cost basis of open positions
		       L string `json:"e"`  // current floating valuation of open positions
		       H string `json:"mf"` // equity = trade balance + unrealized net profit/loss
		       O string `json:"ml"` // free margin = equity - initial margin (maximum margin available to open new positions)
	       } `json:"result"`
}

type Bitflyer []struct {
	CurrencyCode string `json:"currency_code"`
	Amount       int    `json:"amount"`
	Available    int    `json:"available"`
}

type Poloniex struct {
	BTC string `json:"BTC"`
	LTC string `json:"LTC"`
}

type AutoGenerated struct {
	Balance []struct {
		AccountID   string `json:"account_id"`
		Asset       string `json:"asset"`
		Balance     string `json:"balance"`
		Reserved    string `json:"reserved"`
		Unconfirmed string `json:"unconfirmed"`
		Name        string `json:"name"`
	} `json:"balance"`
}

