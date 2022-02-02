package main

import (
	"fmt"
	"reflect"
)

func first() {
	// var x float64 = 3.4
	// v := reflect.ValueOf(x)
	// fmt.Println("type:", v.Type())
	// fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	// fmt.Println("value:", v.Float())

	// var x1 uint8 = 'x'
	// v1 := reflect.ValueOf(x1)
	// fmt.Println("type:", v1.Type())                            // uint8.
	// fmt.Println("kind is uint8: ", v1.Kind() == reflect.Uint8) // true.
	// x1 = uint8(v1.Uint())                                      // v.Uint returns a uint64.
	// fmt.Println(x1)

	type MyInt int
	var x2 MyInt = 7

	v2 := reflect.ValueOf(x2)
	fmt.Println("type:", v2.Type())
	fmt.Println("kind is Int:", v2.Kind() == reflect.Int)
	fmt.Println("Kind: ", v2.Kind())
	fmt.Println("value:", v2.Int())

	type MyInt2 MyInt
	var x3 MyInt2 = 10
	v3 := reflect.ValueOf(x3)
	fmt.Println("type:", v3.Type())
	fmt.Println("kind is Int:", v3.Kind() == reflect.Int)
	fmt.Println("Kind: ", v3.Kind())
	fmt.Println("value:", v3.Int())
}

func printValue(i interface{}) string {
	s := ""
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Int:
		s = fmt.Sprintf("____%d____", v.Int())
	case reflect.Float32:
	case reflect.Float64:
		s = fmt.Sprintf("~~~%7.1e~~~", v.Float())
	case reflect.String:
		s = fmt.Sprintf("```%s```", v.String())
	default:
		s = fmt.Sprintf("%v (%v): %v", v.Type(), v.Kind(), v.Interface())
	}
	return s
}
