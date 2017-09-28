package app

import (
	"github.com/asaskevich/govalidator"
)

func (a *App)setValidators() {

	govalidator.CustomTypeTagMap.Set("validSymbol", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validOrderSymbol", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {


		return false
	}))

	govalidator.CustomTypeTagMap.Set("validOrderAmount", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validOrderType", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validOCOBuy", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validOCOSell", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validOrderId", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validOrderIds", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validOfferCurrency", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validOfferAmount", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validOfferRate", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validOfferPeriod", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validOfferId", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validOrderbookSymbol", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validTickerSymbol", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validStatsSymbol", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validTradesSymbol", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validAssetSymbol", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validPairSymbol", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))

	govalidator.CustomTypeTagMap.Set("validDetailedPairSymbol", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {

		return false
	}))
}

func (p *ParseCore) Validate() (ParseCore, error){
      result, err := govalidator.ValidateStruct(p)
      if err != nil {
	return &ParseCore{}, err
      }
      return result, nil
}
