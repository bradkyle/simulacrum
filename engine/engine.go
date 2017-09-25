package engine

type Engine interface{

}

type MatchingEngineType int
const(
	PriceTimePriority MatchingEngineType = iota
)

