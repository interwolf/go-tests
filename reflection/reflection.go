package main

import (
	"fmt"
	"reflect"
)

// Walk takes a struct x and calls fn for all strings found inside
func Walk(x interface{}, fn func(string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			Walk(val.Index(i).Interface(), fn)
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			Walk(val.Field(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

func main() {
	fmt.Println("Hello Go")
}
