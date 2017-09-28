package app

import (
	"net/http"
	"github.com/thorad/simulacrum/goresponder"
	"github.com/thorad/simulacrum/engine"
)


func (a *App) setResponders(){
	goresponder.SetDefaultResponseFormat("json")
	goresponder.SetLogLevel(goresponder.Verbose)
	goresponder.SetResponseWriter(http.ResponseWriter{})

	//todo goresponder.CustomFormatter.Set([]interface, formatter wrapper)

	goresponder.CustomTypeTagMap.Set(NewOrderParser{}, goresponder.CustomTypeResponder(func(i interface{}, context interface{}) bool {

		return false
	}))

	goresponder.CustomTypeTagMap.Set([]engine.Order{}, goresponder.CustomTypeResponder(func(i interface{}, context interface{}) bool {

		return false
	}))

	//todo goresponder.SetCloser()

}

func (a *App) Respond(w http.ResponseWriter, response interface{}){
	internal_error := goresponder.Respond(response, w)
	if internal_error != nil{
		panic(internal_error)
	}
}