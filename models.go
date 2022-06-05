package main

import (
	"reflect"

	"github.com/SevereCloud/vksdk/marusia"
)

type myPayload struct {
	Index int
	marusia.DefaultPayload
}

type Categories struct {
	Sport    int
	Security int
	Design   int
	Bot      int
	BackEnd  int
	JS       int
	XR       int
	Mobile   int
	MiniApps int
	Marusia  int
}

func (c Categories) MaxCategory() string {
	max := 0
	var name string
	val := reflect.ValueOf(c)
	typeof := val.Type()

	for i := 0; i < val.NumField(); i++ {
		if int(val.Field(i).Int()) > max {
			name = typeof.Field(i).Name
			max = int(val.Field(i).Int())
		}
		// typeOfS.Field(i).Name
		// v.Field(i).Interface()

	}
	if max == 0 {
		name = "XD"
	}
	return name
}
