package parse

type BittrexPairs struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		MarketCurrency     string  `json:"MarketCurrency"`
		BaseCurrency       string  `json:"BaseCurrency"`
		MarketCurrencyLong string  `json:"MarketCurrencyLong"`
		BaseCurrencyLong   string  `json:"BaseCurrencyLong"`
		MinTradeSize       float64 `json:"MinTradeSize"`
		MarketName         string  `json:"MarketName"`
		IsActive           bool    `json:"IsActive"`
		Created            string  `json:"Created"`
	} `json:"result"`
}
