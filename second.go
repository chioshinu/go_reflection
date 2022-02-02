package main

import (
	"fmt"
	"reflect"
)

func second() {
	var x float64 = 3.6
	v := reflect.ValueOf(x)
	y := v.Interface().(float64)
	fmt.Println(y)

	fmt.Printf("value is %7.1e\n", v.Interface())
}
