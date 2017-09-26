package parse

type BittrexAssets struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		Currency        string      `json:"Currency"`
		CurrencyLong    string      `json:"CurrencyLong"`
		MinConfirmation int         `json:"MinConfirmation"`
		TxFee           float64     `json:"TxFee"`
		IsActive        bool        `json:"IsActive"`
		CoinType        string      `json:"CoinType"`
		BaseAddress     interface{} `json:"BaseAddress"`
	} `json:"result"`
}