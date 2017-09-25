package engine

type Orderbook struct{
	Engine 				MatchingEngineType
	Pair				Pair
	BuyQueue			*Orderqueue
	SellQueue			*Orderqueue
}
