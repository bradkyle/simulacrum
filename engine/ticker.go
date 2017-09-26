package engine

import (
	"time"
	"google.golang.org/genproto/googleapis/type/date"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Ticker struct {

}

type Tick struct{
	Interval			int64
	Id				int64

	Symbol				string
	BaseSymbol			string
	QuoteSymbol			string

	DayOpeningPrice			float64
	DayClosingPrice			float64

	OpenPrice			float64
	ClosePrice			float64
	LastPrice			float64
	MidPrice			float64					// bid + ask / 2
	LowPrice			float64					// Lowest trade price of the last 24 hours
	HighPrice			float64

	LowestAsk			float64
	LowestAskSize			float64
	HighestAsk			float64
	HigestAskSize			float64
	LowestBid			float64
	LowestBidSize			float64
	HighestBid			float64
	HighestBidSize			float64

	TotalAskDepth			int64
	TotalBidDepth			int64

	VolumeDay      			float64
	VolumeWeek			float64
	VolumeWeightedAveragePrice	float64
	VolumeBase			float64
	VolumeQuote			float64


	NumberOfTrades			int64

	Time				time.Time
	Date				date.Date
	Timestamp			timestamp.Timestamp
}
