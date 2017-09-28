package app

import (
	"github.com/gorilla/schema"
	"net/http"
	"log"
)

type ParseCore struct{

}

//account
type NewAccountParser struct{
	ParseCore

}

type GetAccountParser struct{
	ParseCore
}

type WithdrawParser struct{
	ParseCore
	WithdrawType	string		`json:"withdraw_type"`				// REQUIRED: has to be one from allowed withdrawal currencies
	Wallet		string		`json:"walletselected"`				// REQUIRED: The wallet to withdraw from, can be “trading”, “exchange”, or “deposit”.
	Amount		float64		`json:"amount"`					// REQUIRED: Amount to withdraw.
	FromAddress	float64		`json:"from_address"`				//
	ToAddress	float64		`json:"to_address"`				//
}

type GetBalanceParser struct{
	ParseCore
	Asset		string          `json:"asset"`	          			// REQUIRED: Asset balance
	Account		string		`json:"account"`				// Account id
}

type TransferParser struct{
	ParseCore
	Amount		float64		`json:"amount"`					// REQUIRED: Amount to transfer.
	Currency	string          `json:"currency"`				// REQUIRED: Currency of funds to transfer.
	WalletFrom	string          `json:"wallet_from"`				// REQUIRED: Wallet to transfer from.
	WalletTo	string          `json:"wallet_to"`				// REQUIRED: Wallet to transfer to.
}

type GetTransferralParser struct{
	ParseCore
}

//orders
type NewOrderParser struct{
	ParseCore
	Symbol		string		`json:"symbol"`					// REQUIRED: and is present
	Amount		float64		`json:"amount"`					// REQUIRED:
	Rate		float64		`json:"price"`					// REQUIRED: must be positive and not null
	Side		string          `json:"side"`					// REQUIRED: buy or sell
	Type		string          `json:"type"`					// REQUIRED: Must be one of allowed order types
	IsHidden	bool            `json:"is_hidden"`				// true if the order should be hidden.
	IsPostonly	bool		`json:"is_postonly"`				// true if the order should be post only. Only relevant for limit orders.
	UseAllAvailable	bool		`json:"use_all_available"`      		// true will post an order that will use all of your available balance.
	OCOOrder        bool 		`json:"ocoorder"`               		// If exchange order: REQUIRED: Set an additional STOP OCO order that will be linked with the current order.
	OCOBuyPrice     float64         `json:"buy_price_oco"`          		// If exchange order: REQUIRED: If exchange order: If ocoorder is true, this field represent the price of the OCO stop order to place
	OCOSellPrice    float64         `json:"sell_price_oco"`         		// If exchange order: REQUIRED: If exchange order: If ocoorder is true, this field represent the price of the OCO stop order to place
}

type CancelOrderParser struct{
	ParseCore
	OrderId		int64		`json:"order_id"`				// REQUIRED: The order ID given by /order/new
}

type CancelMultipleOrdersParser struct{
	ParseCore
	OrderId		[]int64		`json:"order_ids"`				// REQUIRED: An array of the order IDs given by /order/new or /order/new/multi.
}

type GetOrderParser struct{
	ParseCore
	OrderId		[]int64		`json:"order_id"`				// REQUIRED: The order ID given by /order/new.
}

type GetActiveOrdersParser struct{
	ParseCore
	Limit		[]int64		`json:"limit"`					// Limit number of results
}

//offers
type NewOfferParser struct{
	ParseCore
	Currency	string		`json:"currency"`				// REQUIRED: The name of the currency.
	Amount		float64		`json:"amount"`					// REQUIRED: Order size: how much to lend or borrow.
	Rate		float64		`json:"rate"`					// REQUIRED: Rate to lend or borrow at. In percentage per 365 days. (Set to 0 for FRR).
	Period		int32		`json:"period"`					// REQUIRED: Number of days of the funding contract (in days)
	Direction	string		`json:"direction"`				// REQUIRED: Either “lend” or “loan”.
}

type CancelOfferParser struct{
	ParseCore
	OfferId		int32		`json:"offer_id"`				// REQUIRED: The offer ID given by /offer/new.
}

type GetOfferParser struct{
	ParseCore
	OfferId		int32		`json:"offer_id"`				// REQUIRED: The offer ID given by /offer/new.
}

//public
type GetOrderbookParser struct{
	ParseCore
	Symbol		string		`json:"symbol"`					// REQUIRED: The symbol you want information about.
}

type GetTickerParser struct{
	ParseCore
	Symbol		string		`json:"symbol"`					// REQUIRED: The symbol you want information about.
}

type GetStatParser struct{
	ParseCore
	Symbol		string		`json:"symbol"`					// REQUIRED: The symbol you want information about.
}

type GetTradeHistoryParser struct{
	ParseCore
	Symbol		string		`json:"symbol"`					// REQUIRED: The symbol you want information about.
}

type GetAssetParser struct{
	ParseCore
	Symbol		string		`json:"symbol"`					// REQUIRED: The asset symbol you want information about.
}
                                                                        		
type GetPairParser struct{
	ParseCore
	Symbol		string		`json:"symbol"`					// REQUIRED: The symbol you want information about.
}

type GetDetailedPairParser struct{
	ParseCore
	Symbol		string		`json:"symbol",valid:""`			// REQUIRED: The symbol you want information about.
}


type Parser interface{
	Parse(w http.ResponseWriter, r *http.Request)
	Validate(w http.ResponseWriter, r *http.Request)
}

func (p *ParseCore) Parse(w http.ResponseWriter, r *http.Request) *ParseCore{
	r.ParseForm()
	decoder := schema.NewDecoder()
	err := decoder.Decode(p, r.Form)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Status 400"))
		return
	}
	result := p.Validate()
	return result
}

