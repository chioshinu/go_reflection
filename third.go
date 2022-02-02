package main

import (
	"fmt"
	"reflect"
)

func third() {
	var x float64 = 3.4
	v := reflect.ValueOf(x) // take copy of variable
	// v.SetFloat(7.1)         // Error: will panic.
	fmt.Println("settability of v:", v.CanSet())

	var x1 float64 = 3.4
	p1 := reflect.ValueOf(&x1) // Note: take the address of x.
	fmt.Println("type of p1:", p1.Type())
	fmt.Println("settability of p1:", p1.CanSet())

	v1 := p1.Elem()
	fmt.Println("settability of v1:", v1.CanSet())

	v1.SetFloat(7.1)
	fmt.Println(v1.Interface())
	fmt.Println(x1)
	fmt.Println(p1.Interface().(*float64))
}

func structures() {
	type T struct {
		A int `my_tag:"tag_value",json:"_A,omitempty"`
		B string
	}
	t := T{23, "Hello"}
	fmt.Println("t on start ", t)
	s := reflect.ValueOf(&t).Elem()

	typeOfT := s.Type()
	fmt.Printf("%v\n", typeOfT)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v [%v]\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface(), typeOfT.Field(i).Tag)
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Buy!")
	fmt.Println("t is now", t)

	fmt.Println(printValue(10))
	fmt.Println(printValue(10.3))
	fmt.Println(printValue("10"))
	fmt.Println(printValue(t))

}
