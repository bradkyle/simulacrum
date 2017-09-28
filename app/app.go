package app

import (
	"github.com/gorilla/mux"
	"github.com/thorad/simulacrum/engine"
	"log"
	"github.com/thorad/simulacrum/account"
	"net/http"
)

type App struct{
	Router				*mux.Router	
	Config				*Config
	ConfigFile			string
	shutdown 			chan bool
	Accounts			map[string]*account.Account
	Engine				engine.Engine
}

func (app *App)Init() {

	app.loadConfig()
	app.setResponders()
	app.setValidators()

	app.Accounts = make(map[string]*account.Account)


	app.Router = mux.NewRouter().StrictSlash(true)
	app.Router.HandleFunc("/account/new", app.NewAccountHandler).Methods("POST")
	app.Router.HandleFunc("/{exchange}/order/new", app.CreateOrderHandler).Methods("POST")
	app.Router.HandleFunc("/{exchange}/order/cancel", app.CancelOrderHandler).Methods("POST")
	app.Router.HandleFunc("/{exchange}/order", app.OrderHandler).Methods("POST")
	app.Router.HandleFunc("/{exchange}/orders", app.OrdersHandler).Methods("POST")
	app.Router.HandleFunc("/{exchange}/accounts", app.AccountsHandler).Methods("POST")
	app.Router.HandleFunc("/{exchange}/withdraw", app.WithdrawHandler).Methods("POST")
	app.Router.HandleFunc("/{exchange}/transfer", app.TransferHandler).Methods("POST")
	app.Router.HandleFunc("/{exchange}/orderbook", app.OrderbookHandler).Methods("GET")
	app.Router.HandleFunc("/{exchange}/ticker", app.TickerHandler).Methods("GET")
	app.Router.HandleFunc("/{exchange}/trades", app.TradesHandler).Methods("GET")
	app.Router.HandleFunc("/{exchange}/assets", app.AssetsHandler).Methods("GET")
	app.Router.HandleFunc("/{exchange}/pairs", app.AssetsHandler).Methods("GET")
	app.Router.HandleFunc("/bank/account", app.BankHandler).Methods("POST")
}

func (app *App) Start(){
	app.HandleInterrupt()
	app.AdjustGoMaxProcs()
	app.Init()
	app.Start()
	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":6600", app.Router))
	app.Shutdown()
}




