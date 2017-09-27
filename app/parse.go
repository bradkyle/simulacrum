package app

import (
	"github.com/gorilla/schema"
	"net/http"
	"log"
)

type ParseCore struct{

}

type NewAccountParser struct{
	ParseCore
}

type NewOrderParser struct{
	ParseCore
}

type CancelOrderParser struct{
	ParseCore
}

type GetOrderParser struct{
	ParseCore
}

type NewOfferParser struct{
	ParseCore
}

type CancelOfferParser struct{
	ParseCore
}

type GetOfferParser struct{
	ParseCore
}

type GetAccountParser struct{
	ParseCore
}

type GetWithdrawalParser struct{
	ParseCore
}

type WithdrawParser struct{
	ParseCore
}

type GetTransferralParser struct{
	ParseCore
}

type TransferParser struct{
	ParseCore
}

type GetOrderbookParser struct{
	ParseCore
}

type GetTickerParser struct{
	ParseCore
}

type GetAssetParser struct{
	ParseCore
}

type GetPairParser struct{
	ParseCore
}

type GetDetailedPairParser struct{
	ParseCore
}

type Parser interface{
	Validate(w http.ResponseWriter, r *http.Request)
}

func (p *ParseCore) Parse(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	decoder := schema.NewDecoder()
	err := decoder.Decode(p, r.Form)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Status 400"))
		return
	}
}

func (p *ParseCore) Validate(){

}