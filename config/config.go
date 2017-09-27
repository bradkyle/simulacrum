package config

type Config struct{
	Name 				string
	Verbose				bool
	EntropyRate			bool

	TickerInterval			int16

	PairFormat			string
	TimeFormat			string

	RandomAgents			bool
	RandomAgentCount		int

	CurrencyBased 			bool
	CommodityBased  		bool
	OptionsBased			bool
	BondsBased			bool
	SharesBased			bool
	CFDBased			bool

	Ticker				bool
	Book				bool
	Fees				bool
	TradeHistory     		bool
	Margin				bool
	BrokerMargin			bool
	Pairs				bool
	Assets				bool
	HideOrders			bool
	Positions			bool

	ImmediateOrCancel		bool
	PostOnly			bool
	StopOrders			bool
	TrailingStop			bool
	ExpiringOrders			bool

	MaxOrdersPerDay			int64
	MaxOpenOrders			int64

	MaxOrderSize			float64
	MinimumOrderSize		float64
	MaxHistoricTrades		int32

	TakerFeeEnabled			bool
	MakerFeeEnabled			bool
	TransactionFeeEnabled		bool
	WithdrawalFeeEnabled		bool

	TakerFee			float64
	MakerFee			float64
	TransactionFee			float64
	WithdrawalFee			float64

	AssetDelimiter			string						// i.e. -, _ or
	AssetUpperCase			bool						// i.e. usd or USD
	EnabledAssets			[]string
	EnabledBaseAssets		[]string					// i.e. USD
	WithdrawableAssets		[]string
	EnabledAssetPairs		[]Pair

	MatchingEngine			MatchingEngineType

	Lag				int8
	AgentRequestLimit		int16
	OrderFulfilmentLag		int16

	MaxAgentAccounts		int
	InitialTradingBalance		float64
	MinimumSolvencyBalance		float64
	MaximumLeverageAmount		float64
	MaximumTotalDebt		float64
	InitialMargin			float64

	AllowedAccountTypes		[]BalanceType
	AllowedOrderExecutions		[]OrderExecution
	AllowedAssetTypes		[]AssetType

	WithdrawEnabled			bool
	TransfersEnabled		bool
	WithdrawLag			int64
	TransferLag			int64

	OrderNotProcessed		bool
	ErrorOnRequest			bool
	PlaceWrongOrder			bool
	InacurateInformation		bool
	UnauthorizedAcess		bool

	MaximumEpisodeLength		int32

}
