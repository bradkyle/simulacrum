package asset


type Pair struct{
	Id				int64
	Symbol				string
	Base				Asset
	Quote				Asset
	MinimumOrderSize		float64
	MaximumOrderSize		float64
	MinimumOrderIncrement		float64
	MinimumPriceIncrement		float64
	MinimumMargin			float64
	InitMargin			float64
	Enabled				bool
}
