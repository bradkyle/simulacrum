package parse

type BitfinexMarginInformation []struct {
	MarginBalance     string `json:"margin_balance"`
	TradableBalance   string `json:"tradable_balance"`
	UnrealizedPl      string `json:"unrealized_pl"`
	UnrealizedSwap    string `json:"unrealized_swap"`
	NetValue          string `json:"net_value"`
	RequiredMargin    string `json:"required_margin"`
	Leverage          string `json:"leverage"`
	MarginRequirement string `json:"margin_requirement"`
	MarginLimits      []struct {
		OnPair            string `json:"on_pair"`
		InitialMargin     string `json:"initial_margin"`
		MarginRequirement string `json:"margin_requirement"`
		TradableBalance   string `json:"tradable_balance"`
	} `json:"margin_limits"`
	Message string `json:"message"`
}

type PoloniexMarginAccountSummary struct {
	TotalValue         string `json:"totalValue"`
	Pl                 string `json:"pl"`
	LendingFees        string `json:"lendingFees"`
	NetValue           string `json:"netValue"`
	TotalBorrowedValue string `json:"totalBorrowedValue"`
	CurrentMargin      string `json:"currentMargin"`
}