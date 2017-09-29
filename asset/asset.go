package asset

type AssetType int
const(
	Cryptocurrency	AssetType = iota
	Fiat
	Bond
	Commodity
	Share
	CTFO
)

type Asset struct{
	Name				string
	Symbol				string
	AssetType			AssetType
	Enabled				bool
	MaxOrderSize			float64
	MinOrderSize			float64
}



