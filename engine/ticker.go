package engine

type Ticker struct {

}

type Tick struct{
	Interval			int64
	Symbol				string
	BaseSymbol			string
	QuoteSymbol			string
	LastPrice			float64
	MidPrice			float64					// bid + ask / 2
	LowPrice			float64					// Lowest trade price of the last 24 hours
	HighPrice			float64
	LowestAsk			float64
	HighestAsk			float64
	Volume        			float64
	VolumeWeightedAveragePrice	float64
	NumberOfTrades			float64
	DayOpeningPrice			float64
	ValidTime			float64
	VolumeBase			float64
	VolumeQuote			float64
}
