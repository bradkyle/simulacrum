package goresponder

import "sync"

type CustomTypeResponder func(i interface{}, o interface{}) bool

type Responder func(str string) bool

type customTypeTagMap struct {
	responders map[interface{}]CustomTypeResponder

	sync.RWMutex
}

func (tm *customTypeTagMap) Get(kind interface{}) (CustomTypeResponder, bool) {
	tm.RLock()
	defer tm.RUnlock()
	v, ok := tm.responders[kind]
	return v, ok
}

func (tm *customTypeTagMap) Set(kind interface{}, ctv CustomTypeResponder) {
	tm.Lock()
	defer tm.Unlock()
	tm.responders[kind] = ctv
}

// CustomTypeTagMap is a map of functions that can be used as tags for ValidateStruct function.
// Use this to validate compound or custom types that need to be handled as a whole, e.g.
// `type UUID [16]byte` (this would be handled as an array of bytes).
var CustomTypeTagMap = &customTypeTagMap{responders: make(map[interface{}]CustomTypeResponder)}

var ResponseMap = map[interface{}]Responder{
	error: 		RespondError,
}
