package swiftVisitor

import (
	"fmt"
	"reflect"
)

var VOID = &SwiftValue{}
var NULL = &SwiftValue{}
var INVALID = &SwiftValue{}
var RETURNVOID = &SwiftValue{}
var BREAK = &SwiftValue{}
var ERROR = &SwiftValue{}
var CONTINUE = &SwiftValue{}
var VECTO = &SwiftValue{}

type SwiftValue struct {
	value interface{}
}

func (v *SwiftValue) isInt() bool {
	return reflect.TypeOf(v.value).Kind() == reflect.Int
}

func (v *SwiftValue) asInt() int {
	if v.isDouble() {
		return int(v.asDouble())
	}
	return v.value.(int)
}

func (v *SwiftValue) isDouble() bool {
	return reflect.TypeOf(v.value).Kind() == reflect.Float64
}

func (v *SwiftValue) asDouble() float64 {
	if v.isInt() {
		return float64(v.asInt())
	}
	return v.value.(float64)
}

func (v *SwiftValue) isNumber() bool {
	return v.isInt() || v.isDouble()
}

func (v *SwiftValue) isString() bool {
	return reflect.TypeOf(v.value).Kind() == reflect.String
}

func (v *SwiftValue) asString() string {
	return v.value.(string)
}
func (v *SwiftValue) isChar() bool {
	a:=false
	if len(v.asString())>2{
		a=true
	}
	return a
}

func (v *SwiftValue) asChar() string {
	return v.value.(string)
}

func (v *SwiftValue) isBool() bool {
	return reflect.TypeOf(v.value).Kind() == reflect.Bool
}

func (v *SwiftValue) asBool() bool {
	return v.value.(bool)
}

func (v *SwiftValue) isList() bool {
	return reflect.TypeOf(v.value).Kind() == reflect.Slice
}

func (v *SwiftValue) asList() []*SwiftValue {
	return v.value.([]*SwiftValue)
}

func (v *SwiftValue) equals(other *SwiftValue) bool {
	if v == NULL {
		return other == NULL
	}
	if v == other {
		return true
	}
	if v.isInt() && other.isInt() {
		return v.asInt() == other.asInt()
	}
	if v.isDouble() && other.isDouble() {
		return v.asDouble() == other.asDouble()
	}
	if v.isString() && other.isString() {
		return v.asString() == other.asString()
	}
	if v.isBool() && other.isBool() {
		return v.asBool() == other.asBool()
	}
	return false
}

func (v *SwiftValue) String() string {
	if v.isList() {
		arr := v.asList()
		res := "["
		for i := 0; i < len(arr)-1; i += 1 {
			res += arr[i].String() + ", "
		}
		res += arr[len(arr)-1].String() + "]"
		return res
	} else {
		return fmt.Sprintf("%v", v.value)
	}
}