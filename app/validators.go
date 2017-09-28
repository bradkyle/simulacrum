package app

import (
	"github.com/asaskevich/govalidator"
	"log"
)

func (a *App)SetValidators() {
	govalidator.CustomTypeTagMap.Set("customByteArrayValidator", CustomTypeValidator(func(i interface{}, context interface{}) bool {
		switch v := context.(type) { // you can type switch on the context interface being validated
		case StructWithCustomByteArray:
		// you can check and validate against some other field in the context,
		// return early or not validate against the context at all – your choice
		case SomeOtherType:
		// ...
		default:
		// expecting some other type? Throw/panic here or continue
		}

		switch v := i.(type) { // type switch on the struct field being validated
		case CustomByteArray:
			for _, e := range v {
				// this validator checks that the byte array is not empty, i.e. not all zeroes
				if e != 0 {
					return true
				}
			}
		}
		return false
	}))
}

func (p *ParseCore) Validate() ParseCore{
      result, err := govalidator.ValidateStruct(p)
      if err != nil {
	log.Println(err)
      }
      return result
}
