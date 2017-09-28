package goresponder

import (
	"reflect"
	"net/http"
	"log"
	"errors"
	"fmt"
	"encoding/json"
)

var (
	responseFormat 	bool
	responseWriter  interface{}
	logLevel	logLevel
)


type logLevel int
const(
	Verbose	logLevel = iota

)

func SetDefaultResponseFormat(value string) {
	if value == "json" || "csv" || "xml" || "protobuf" || "bson"{
		responseFormat = value
	} else{
		panic("The format provided to goresponder is invalid!")
	}
}

func SetDefaultResponder(value string) {

}

func SetLogLevel(level logLevel) {

}

//todo set closer will enforce common functionality for deleting data after a responder has been called
func SetCloser(){
	print("hello")
}

//todo 	goresponder.SetResponseWriter()
func SetResponseWriter(w interface{}){
	print("hello")
}

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

func DefaultResponder(input struct{}, w http.ResponseWriter){
	if logLevel == Verbose{
		log.Println("Using default responder...")
	}

	if responseFormat == "json"{
		b, err := json.Marshal(input)
		if err != nil {
			panic(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}

	return
}