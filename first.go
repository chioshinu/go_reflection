package main

import (
	"fmt"
	"reflect"
)

func first() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	var x1 uint8 = 'x'
	v1 := reflect.ValueOf(x1)
	fmt.Println("type:", v1.Type())                            // uint8.
	fmt.Println("kind is uint8: ", v1.Kind() == reflect.Uint8) // true.
	x1 = uint8(v1.Uint())                                      // v.Uint returns a uint64.
	fmt.Println(x1)

	type MyInt int
	var x2 MyInt = 7
	v2 := reflect.ValueOf(x2)
	fmt.Println("type:", v2.Type())
	fmt.Println("kind is float64:", v2.Kind() == reflect.Float64)
	fmt.Println("value:", v2.Int())
}
