package app

import (
	"net/http"
	"github.com/thorad/simulacrum/goresponder"
)


func (a *App) setResponders(){
	goresponder.SetResponseFormat("json")

	goresponder.CustomTypeTagMap.Set("validOrderbookSymbol", goresponder.CustomTypeResponder(func(i interface{}, context interface{}) bool {

		return false
	}))

	goresponder.CustomTypeTagMap.Set("validOrderbookSymbol", goresponder.CustomTypeResponder(func(i interface{}, context interface{}) bool {

		return false
	}))

}

func (a *App) Respond(w http.ResponseWriter, response interface{}, err error){
	if err != nil {
		internal_error := goresponder.Respond(err, w)
		if internal_error != nil{
			panic(internal_error)
		}
		return
	}
	internal_error := goresponder.Respond(response, w)
	if internal_error != nil{
		panic(internal_error)
	}
}