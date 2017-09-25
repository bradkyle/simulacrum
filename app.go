package main

import (
	"github.com/gorilla/mux"
	"github.com/thorad/simulacrum/config"
	"github.com/thorad/simulacrum/common"
	"github.com/thorad/deeptrade/work/simex/account"
	"github.com/thorad/simulacrum/engine"
	"flag"
	"log"
	"os"
	"runtime"
	"strconv"
	"os/signal"
	"syscall"
	"fmt"
	"net/http"
	"errors"
	"encoding/json"
)

type App struct{
	Router				*mux.Router	
	Config				*config.Config
	ConfigFile			string
	Logger				*common.Logger
	
	shutdown 			chan bool
	Accounts			map[string]*account.Account
	Orderbooks			map[string]*engine.Orderbook
	Tickers				map[string]*engine.Ticker
	Engine				engine.Engine
}

func (app *App)Init() {
	app.Accounts = make(map[string]*account.Account)
	app.Orderbooks = make(map[string]*engine.Orderbook)
	app.Tickers = make(map[string]*engine.Ticker)
	
	flag.StringVar(&app.ConfigFile, "config", config.GetFilePath(""), "config file to load")
	flag.Parse()

	app.Config = new(config.Config)
	log.Printf("Loading config file %s..\n", app.ConfigFile)

	err := app.Config.LoadConfig(app.ConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	app.Load
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
	app.Router.HandleFunc("/{exchange}/trades", app.TradeHistoryHandler).Methods("GET")
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

func (a *App) FindAccountByScope(scope string) (*account.Account, error) {
	accounts := a.Accounts
	if scope != "" {
		if ag, ok := accounts[scope]; ok {
			return ag, nil
		}
		return &account.Account{}, errors.New("")
	}
	return &account.Account{}, errors.New("")
}

func (a *App) Authenticate(w http.ResponseWriter, r *http.Request) (){
	ag, err := a.FindAccountByScope(scope)
	if err != nil{
		return &agent.Agent{}, &exchange.Exchange{}, err
	}
	if err != nil{
		log.Println("CreateOrderandler -> HandleRequest: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Status 400"))
		return
	}

}

// New Primary Entity Handlers
func (a *App) NewAccountHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ag := new(account.Account)
	decoder := schema.NewDecoder()
	err := decoder.Decode(ag, r.Form)
	if err != nil {
		log.Println("NewAccountHandler -> Decode: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Status 400"))
		return
	}

	err, _ = ag.Parse()
	if err != nil {
		log.Println("NewAccountHandler -> Setup: ", err)
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Status 408"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status 200"))
}


// Private Handlers
func (a *App) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	order := new(parse.Order)
	decoder := schema.NewDecoder()
	err = decoder.Decode(order, r.Form)
	if err != nil {
		log.Println("CreateOrderHandler -> Decode: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Status 400"))
		return
	}

	err = engine.AddOrder(order, ag)
	if err != nil {
		log.Println("CreateOrderandler -> AddOrder: ", err)
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Status 408"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Status 201"))
}

func (a *App) CancelOrderHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	order := new(orderbook.Order)
	decoder := schema.NewDecoder()
	err = decoder.Decode(order, r.Form)
	if err != nil {
		log.Println("CancelOrderandler -> Decode: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Status 400"))
		return
	}

	err = engine.CancelOrder(order, ag)
	if err != nil {
		log.Println("CancelOrderHandler -> CancelOrder: ", err)
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Status 408"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status 200"))
}

func (a *App) OrderHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	order := new(orderbook.Order)
	decoder := schema.NewDecoder()
	err = decoder.Decode(order, r.Form)
	if err != nil {
		log.Println("OrderHandler -> Decode: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Status 400"))
		return
	}

	return_order, err := engine.ReturnOrder(order, ag)
	if err != nil {
		log.Println("OrderHandler -> ReturnOrder: ", err)
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Status 408"))
		return
	}

	json.NewEncoder(w).Encode(return_order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status 200"))
}

func (a *App) OrdersHandler(w http.ResponseWriter, r *http.Request) {
	return_orders, err := engine.ReturnOrders(ag)
	if err != nil {
		log.Println("OrdersHandler -> ReturnOrders: ", err)
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Status 408"))
		return
	}

	json.NewEncoder(w).Encode(return_orders)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status 200"))
}

func (a *App) AccountsHandler(w http.ResponseWriter, r *http.Request) {
	return_accounts, err := engine.ReturnAccounts(ag)
	if err != nil {
		log.Println("OrdersHandler -> ReturnOrders: ", err)
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Status 408"))
		return
	}

	json.NewEncoder(w).Encode(return_accounts)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status 200"))
}

func (a *App) WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	withdrawal := new(account.Withdrawl)
	decoder := schema.NewDecoder()
	err = decoder.Decode(withdrawal, r.Form)
	if err != nil {
		log.Println("OrderHandler -> Decode: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Status 400"))
		return
	}

	withdrawal, err = engine.Withdraw(withdrawal,ag)
	if err != nil {
		log.Println("OrdersHandler -> ReturnOrders: ", err)
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Status 408"))
		return
	}

	fmt.Println(withdrawal)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status 200"))
}

func (a *App) TransferHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	transferal := new(account.Transferal)
	decoder := schema.NewDecoder()
	err = decoder.Decode(transferal, r.Form)
	if err != nil {
		log.Println("OrderHandler -> Decode: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Status 400"))
		return
	}

	transferal, err = engine.Transfer(transferal,ag)
	if err != nil {
		log.Println("OrdersHandler -> ReturnOrders: ", err)
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Status 408"))
		return
	}

	fmt.Println(transferal)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status 200"))
}

// Public Handlers

func (a *App) OrderbookHandler(w http.ResponseWriter, r *http.Request) {
	return_orderbooks, err := engine.ReturnOrderbooks()
	if err != nil {
		log.Println("OrderbookHandler -> ReturnOrderbook: ", err)
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Status 408"))
		return
	}

	json.NewEncoder(w).Encode(return_orderbooks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status 200"))
}

func (a *App) TickerHandler(w http.ResponseWriter, r *http.Request) {
	return_tickers, err := engine.ReturnTickers()
	if err != nil {
		log.Println("TickerHandler -> ReturnTickers: ", err)
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Status 408"))
		return
	}

	json.NewEncoder(w).Encode(return_tickers)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status 200"))
}

func (a *App) TradeHistoryHandler(w http.ResponseWriter, r *http.Request) {
	err = engine.ReturnTradeHistory()
	if err != nil {
		log.Println("TradeHistoryHandler -> ReturnTradeHistory: ", err)
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Status 408"))
		return
	}

	//json.NewEncoder(w).Encode(return_tickers)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status 200"))
}

func (a *App) AssetsHandler(w http.ResponseWriter, r *http.Request) {
	return_assets, err := exch.ReturnAssets()
	if err != nil {
		log.Println("AssetsHandler -> ReturnAssets: ", err)
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Status 408"))
		return
	}

	json.NewEncoder(w).Encode(return_assets)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status 200"))
}

func (a *App) PairsHandler(w http.ResponseWriter, r *http.Request) {
	return_assets, err := exch.ReturnPairs()
	if err != nil {
		log.Println("AssetsHandler -> ReturnPairs: ", err)
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Status 408"))
		return
	}

	json.NewEncoder(w).Encode(return_assets)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status 200"))
}

// This acts as a scoreboard component for the simulation essentially
// It only takes base fiat currencies i.e. USD, EUR, GBP .etc
func (a *App) BankHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status 200"))
}

// Debug and Analytics
func (a *App) ActualRewardsHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status 200"))
}


// AdjustGoMaxProcs adjusts the maximum processes that the CPU can handle.
func (a *App) AdjustGoMaxProcs() {
	log.Println("Adjusting bot runtime performance..")
	maxProcsEnv := os.Getenv("GOMAXPROCS")
	maxProcs := runtime.NumCPU()
	log.Println("Number of CPU's detected:", maxProcs)

	if maxProcsEnv != "" {
		log.Println("GOMAXPROCS env =", maxProcsEnv)
		env, err := strconv.Atoi(maxProcsEnv)
		if err != nil {
			log.Println("Unable to convert GOMAXPROCS to int, using", maxProcs)
		} else {
			maxProcs = env
		}
	}
	if i := runtime.GOMAXPROCS(maxProcs); i != maxProcs {
		log.Fatal("Go Max Procs were not set correctly.")
	}
	log.Println("Set GOMAXPROCS to:", maxProcs)
}


// HandleInterrupt monitors and captures the SIGTERM in a new goroutine then
// shuts down bot
func (a *App) HandleInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		log.Printf("Captured %v.", sig)
		a.Shutdown()
	}()
}

// Shutdown correctly shuts down bot saving configuration files
func (a *App) Shutdown() {
	<-a.shutdown
	log.Println("Bot shutting down..")

	// do stuff before shutdown

	log.Println("Exiting.")
	os.Exit(1)
}