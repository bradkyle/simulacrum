package parse

import (
	"github.com/thorad/simulacrum/engine"
	"github.com/thorad/simulacrum/asset"
	"github.com/thorad/simulacrum/config"
	"github.com/thorad/simulacrum/account"
)

type Parser interface {

}

func TickersParser(config config.Config, tickers []*engine.Ticker, tickerparser Parser){

}

func StatsParser(config config.Config, stats []*engine.Stats, statsparser Parser){

}

func FundingbooksParser(config config.Config, fundingbooks []*engine.Fundingbook, fundingbookparser Parser){

}

func OrderbooksParser(config config.Config, orderbooks []*engine.Orderbook, orderbookparser Parser){

}

func PairsParser(config config.Config, pairs []*asset.Pair, pairsparser Parser){

}

func AssetsParser(config config.Config, assets []*asset.Asset, assetsparser Parser){

}

func ErrorsParser(config config.Config, error []*engine.Error, errorsparser Parser){

}

func AccountParser(config config.Config, accounts []*account.Account, accountparser Parser){

}

func OrderParser(config config.Config, ){

}

func OfferParser(config config.Config, ){

}

func TradesParser(config config.Config, trades []*engine.Trade, tradesparser Parser){

}

func LendsParser(config config.Config, lends []*engine.Lend, lendsparser Parser){

}

/*
Bithumb
Bitfinex
coinone
Bittrex
HitBTC
Kraken
Poloniex
Quoine
Bitstamp
bitFlyer
OKCoin
Gemini
Korbit
Liqui
Gatecoin
BTC38
YoBit
LakeBTC
CHBTC
Vaultoro
Binance
EXMO
CEX.IO
itBit
Tidex
Bitcoin
Allcoin
Livecoin
viaBTC
C-Cex
BTCTrade
BitcoinIndonesia
Bter
QuadrigaCX
Cryptopia
Coinfloor
BTCMarkets
BTCC
BitMarket
ACX
BitXSouthAfrica
Bitso
TheRockTrading
CoinMate
CoinExchange
Coinsecure
Novaexchange
VirWox
Bit2c
Bleutrade
Bittylicious
Bisq
Bitkonan
CryptoX
Vircurex
Electrocoin
BitBay
BTC-e
Indacoin
UseCryptos
Cryptonit
Yunbi
*/
