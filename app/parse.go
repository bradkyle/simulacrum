package app

import (
	"github.com/gorilla/schema"
	"net/http"
	"log"
	"errors"
)

// Sets a parser for each request handler

/*
	X-GEMINI-APIKEY	Your Gemini API key
	X-GEMINI-PAYLOAD	The base64-encoded JSON payload
	X-GEMINI-SIGNATURE	hex(HMAC_SHA384(base64(payload), key=api_secret)
*/

// parameters required in every request
// agent scope should be passed with every request for debugging purposes
type ParseCore struct{
	AgentScope	string
	Response	interface{}
}

// parameters required in authenticated endpoints
type ParseAuth struct{

}

//account
type NewAccountParser struct{
	ParseCore
	Index		string
	Name 		string
	Balances	string
}

type GetAccountParser struct{
	ParseCore
}

type WithdrawParser struct{
	ParseCore
	WithdrawType	string		`json:"withdraw_type"`						// REQUIRED: has to be one from allowed withdrawal currencies
	Wallet		string		`json:"walletselected"`						// REQUIRED: The wallet to withdraw from, can be “trading”, “exchange”, or “deposit”.
	Amount		float64		`json:"amount"`							// REQUIRED: Amount to withdraw.
	FromAddress	float64		`json:"from_address"`						//
	ToAddress	float64		`json:"to_address"`						//
}                                                                                               	
                                                                                                	
type GetBalanceParser struct{                                                                   	
	ParseCore                                                                               	
	Symbol		string          `json:"symbol",valid:"validAssetSymbol"`	          	// REQUIRED: Asset balance
	Account		string		`json:"account",valid:"validAccountId"`				// Account id
}                                                                                               	
                                                                                                	
type TransferParser struct{                                                                     	
	ParseCore                                                                               	
	Amount		float64		`json:"amount"`							// REQUIRED: Amount to transfer.
	Currency	string          `json:"currency"`						// REQUIRED: Currency of funds to transfer.
	WalletFrom	string          `json:"wallet_from"`						// REQUIRED: Wallet to transfer from.
	WalletTo	string          `json:"wallet_to"`						// REQUIRED: Wallet to transfer to.
}                                                                                               	
                                                                                                	
type GetTransferralParser struct{                                                               	
	ParseCore                                                                               	
}                                                                                               	
                                                                                                	
//orders                                                                                        	
type NewOrderParser struct{                                                                     	
	ParseCore                                                                               	
	Symbol		string		`json:"symbol",valid:"alphanum,validOrderSymbol"`		// REQUIRED: and is present
	Amount		float64		`json:"amount",valid:"float,validOrderAmount"`			// REQUIRED:
	Rate		float64		`json:"price",valid:"float,validOrderPrice"`			// REQUIRED: must be positive and not null
	Side		string          `json:"side",valid:"alphanum,in(buy|sell)"`			// REQUIRED: buy or sell
	Type		string          `json:"type",valid:"alphanum,validOrderType"`			// REQUIRED: Must be one of allowed order types
	IsHidden	bool            `json:"is_hidden",valid:"in(true|false),optional"`		// true if the order should be hidden.
	IsPostonly	bool		`json:"is_postonly",valid:"in(true|false),optional"`		// true if the order should be post only. Only relevant for limit orders.
	UseAllAvailable	bool		`json:"use_all_available",valid:"in(true|false),optional"`      // true will post an order that will use all of your available balance.
	OCOOrder        bool 		`json:"ocoorder",valid:"in(true|false)"`               		// If exchange order: REQUIRED: Set an additional STOP OCO order that will be linked with the current order.
	OCOBuyPrice     float64         `json:"buy_price_oco",valid:"float,validOCOBuy"`          	// If exchange order: REQUIRED: If exchange order: If ocoorder is true, this field represent the price of the OCO stop order to place
	OCOSellPrice    float64         `json:"sell_price_oco",valid:"float,validOCOSell"`         	// If exchange order: REQUIRED: If exchange order: If ocoorder is true, this field represent the price of the OCO stop order to place
}                                                                                               	
                                                                                                	
type CancelOrderParser struct{                                                                  	
	ParseCore                                                                               	
	OrderId		int64		`json:"order_id",valid:"int,validOrderId"`			// REQUIRED: The order ID given by /order/new
}                                                                                               	
                                                                                                	
type CancelMultipleOrdersParser struct{                                                         	
	ParseCore                                                                               	
	OrderId		[]int64		`json:"order_ids",valid:"int,validOrderIds"`			// REQUIRED: An array of the order IDs given by /order/new or /order/new/multi.
}                                                                                               	
                                                                                                	
type GetOrderParser struct{                                                                     	
	ParseCore                                                                               	
	OrderId		[]int64		`json:"order_id",valid:"int,validOrderId"`			// REQUIRED: The order ID given by /order/new.
}                                                                                               	
                                                                                                	
type GetActiveOrdersParser struct{                                                              	
	ParseCore                                                                               	
	Limit		[]int64		`json:"limit",valid:"int,optional"`				// Limit number of results
}                                                                                               	
                                                                                                	
//offers                                                                                        	
type NewOfferParser struct{                                                                     	
	ParseCore                                                                               	
	Currency	string		`json:"currency",valid:"alphanum,validOfferCurrency"`		// REQUIRED: The name of the currency.
	Amount		float64		`json:"amount",valid:"float,validOfferAmount"`			// REQUIRED: Order size: how much to lend or borrow.
	Rate		float64		`json:"rate",valid:"float,validOfferRate"`			// REQUIRED: Rate to lend or borrow at. In percentage per 365 days. (Set to 0 for FRR).
	Period		int32		`json:"period",valid:"int,validOfferPeriod"`			// REQUIRED: Number of days of the funding contract (in days)
	Direction	string		`json:"direction",valid:"in(lend|loan)"`			// REQUIRED: Either “lend” or “loan”.
}                                                                                               	
                                                                                                	
type CancelOfferParser struct{                                                                  	
	ParseCore                                                                               	
	OfferId		int32		`json:"offer_id",valid:"int,validOfferId"`			// REQUIRED: The offer ID given by /offer/new.
}                                                                                               	
                                                                                                	
type GetOfferParser struct{                                                                     	
	ParseCore                                                                               	
	OfferId		int32		`json:"offer_id",valid:"int,validOfferId"`			// REQUIRED: The offer ID given by /offer/new.
}                                                                                               	
                                                                                                	
//public                                                                                        	
type GetOrderbookParser struct{                                                                 	
	ParseCore                                                                               	
	Symbol		string		`json:"symbol",valid:"validOrderbookSymbol"`			// REQUIRED: The symbol you want information about.
}                                                                                               	
                                                                                                	
type GetTickerParser struct{                                                                    	
	ParseCore                                                                               	
	Symbol		string		`json:"symbol",valid:"validTickerSymbol"`			// REQUIRED: The symbol you want information about.
}                                                                                               	
                                                                                                	
type GetStatParser struct{                                                                      	
	ParseCore                                                                               	
	Symbol		string		`json:"symbol",valid:"validStatsSymbol"`			// REQUIRED: The symbol you want information about.
}                                                                                               	
                                                                                                	
type GetTradeHistoryParser struct{                                                              	
	ParseCore                                                                               	
	Symbol		string		`json:"symbol",valid:"validTradesSymbol"`			// REQUIRED: The symbol you want information about.
}                                                                                               	
                                                                                                	
type GetAssetParser struct{                                                                     	
	ParseCore                                                                               	
	Symbol		string		`json:"symbol",valid:"validAssetSymbol"`			// REQUIRED: The asset symbol you want information about.
}                                                                                               	
                                                                                                	
type GetPairParser struct{                                                                      	
	ParseCore                                                                               	
	Symbol		string		`json:"symbol",valid:"validPairSymbol"`				// REQUIRED: The symbol you want information about.
}                                                                                               	
                                                                                                	
type GetDetailedPairParser struct{                                                              	
	ParseCore                                                                               	
	Symbol		string		`json:"symbol",valid:"validDetailedPairSymbol"`			// REQUIRED: The symbol you want information about.
}

// Endpoint parsers with no input
type GetAccountsParser 		struct{ParseCore}
type GetBalancesParser 		struct{ParseCore}
type GetBalanceParser 		struct{ParseCore}
type GetWithdrawalsParser 	struct{ParseCore}
type GetBalancesParser 		struct{ParseCore}
type GetTransferalsParser 	struct{ParseCore}
type GetBalancesParser 		struct{ParseCore}
type GetTradeHistoriesParser	struct{ParseCore}
type GetFundingbookParser 	struct{ParseCore}
type GetLendsParser 		struct{ParseCore}
type GetOrderbooksParser 	struct{ParseCore}
type GetStatsParser 		struct{ParseCore}
type GetTickersParser 		struct{ParseCore}
type GetAssetsParser 		struct{ParseCore}
type GetPairsParser 		struct{ParseCore}
type GetDetailedPairsParser 	struct{ParseCore}
type CancelAllOrdersParser 	struct{ParseCore}
type GetOrdersParser 		struct{ParseCore}
type GetOffersParser 		struct{ParseCore}
type CancelAllOffersParser 	struct{ParseCore}


func (a *App) Parse(p *ParseCore, w http.ResponseWriter, r *http.Request) *ParseCore{

	p.AgentScope = r.Header

	r.ParseForm()
	decoder := schema.NewDecoder()
	err := decoder.Decode(p, r.Form)
	if err != nil {
		a.Respond(w,errors.New("")) //todo replace with custom error struct and responder
		return &ParseCore{}
	}
	result, err := p.Validate()
	if err != nil {
		a.Respond(w,errors.New("")) //todo replace with custom error struct and responder
		return &ParseCore{}
	}
	return result
}

