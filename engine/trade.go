package engine

import (
	"time"
	"github.com/thorad/simulacrum/account"
	"github.com/thorad/simulacrum/asset"
)

type Trade struct{
	Account                 *account.Account                           //
	Amount                  float64                                    //
	Price                   float64	                                   //
	Side			OrderSide
	Base                    asset.Asset                                //
	Target                  asset.Asset                                //
	OrderExecution          OrderExecution                             //
	Complete                bool                                       //
	TimeCreated             time.Time                                  //
	TimeCompleted           time.Time                                  //
}

type AnonymizedTrade struct {
	Amount int     `json:"shares"`
	Symbol string  `json:"ticker"`
	Price  float64 `json:"price"`
	Kind   string  `json:"kind"`
	Time   int64   `json:"time"`
}

func fill(){

}

//func AnonymizeTrade(trade Trade) AnonymizedTrade {
//	return AnonymizedTrade{
//		Amount: trade.Shares,
//		Ticker: trade.Ticker,
//		Price:  trade.Price,
//		Kind:   trade.Kind,
//		Time:   trade.Time,
//	}
//}

