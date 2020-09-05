package main

import (
	"reflect"
)

type aaa struct {
	A int
}

func main() {
	a := aaa{A: 1}
	setFields(&a, 2, "A")
	println(a.A)
}

func setFields(dst, src interface{}, name string) {
	d := reflect.ValueOf(dst).Elem()
	s := reflect.ValueOf(src)
	df := d.FieldByName(name)
	switch df.Kind() {
	case reflect.String:
		if v := s.String(); v != "" {
			df.SetString(v)
		}
	case reflect.Int:
		if v := s.Int(); v != 0 {
			df.SetInt(v)
		}
	case reflect.Float32:
	case reflect.Float64:
		if v := s.Float(); v != 0 {
			df.SetFloat(v)
		}
	case reflect.Uint:
		if v := s.Uint(); v != 0 {
			df.SetUint(v)
		}
		// handle other kinds
	}
}
