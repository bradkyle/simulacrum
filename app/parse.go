package app

import (
	"net/http"

	"github.com/thorad/simulacrum/parse"
	"github.com/gorilla/schema"
	"errors"
)

func (a *App) Parse(p *parse.ParseCore, w http.ResponseWriter, r *http.Request) *parse.ParseCore{

	p.AgentScope = r.Header

	r.ParseForm()
	decoder := schema.NewDecoder()
	err := decoder.Decode(p, r.Form)
	if err != nil {
		a.Respond(w,errors.New("")) //todo replace with custom error struct and responder
		return &parse.ParseCore{}
	}
	result, err := Validate(p)
	if err != nil {
		a.Respond(w,errors.New("")) //todo replace with custom error struct and responder
		return &parse.ParseCore{}
	}
	return result
}

