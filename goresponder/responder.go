package goresponder

import (
	"reflect"
	"fmt"
	"net/http"
	"log"
	"errors"
)

var (
	responseFormat 	bool
	responseWriter  interface{}
)


func SetResponseFormat(value string) {
	if value == "json" || "csv" || "xml" || "protobuf" || "bson"{
		responseFormat = value
	} else{
		panic("The format provided to goresponder is invalid!")
	}
}

//todo 	goresponder.SetResponseWriter()

func RespondError(err error, w http.ResponseWriter){
	log.Println(err)
	w.WriteHeader(http.StatusRequestTimeout)
	w.Write([]byte("Status 408"))
	return
}

// ValidateStruct use tags for fields.
// result will be equal to `false` if there are any errors.
func Respond(s interface{}, w http.ResponseWriter) error{
	if s == nil {
		return errors.New("There was no type provided to Respond")
	}
	var err error
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	// we only accept structs
	//if val.Kind() != reflect.Struct {
	//	return fmt.Errorf("function only accepts structs; got %s", val.Kind())
	//}

	if valueField.Kind() == reflect.Struct && typeField.Tag.Get(tagName) != "-" {
		var err error
		structResult, err = ValidateStruct(valueField.Interface())
		if err != nil {
			errs = append(errs, err)
		}
	}

	var errs Errors
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		if typeField.PkgPath != "" {
			continue // Private field
		}
		structResult := true
		if valueField.Kind() == reflect.Struct && typeField.Tag.Get(tagName) != "-" {
			var err error
			structResult, err = ValidateStruct(valueField.Interface())
			if err != nil {
				errs = append(errs, err)
			}
		}
		resultField, err2 := typeCheck(valueField, typeField, val, nil)
		if err2 != nil {

			// Replace structure name with JSON name if there is a tag on the variable
			jsonTag := toJSONName(typeField.Tag.Get("json"))
			if jsonTag != "" {
				switch jsonError := err2.(type) {
				case Error:
					jsonError.Name = jsonTag
					err2 = jsonError
				case Errors:
					err2 = jsonError
				}
			}

			errs = append(errs, err2)
		}
		result = result && resultField && structResult
	}
	if len(errs) > 0 {
		err = errs
	}
	return result, err
}
