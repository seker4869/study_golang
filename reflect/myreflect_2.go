package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	filed1 bool   "An import answer"
	filed2 string "The name of the thing"
	filed3 int    "how much there are"
}

func main() {
	tt := TagType{true, "Barak Obama", 1}
	var reflectType reflect.Type = reflect.TypeOf(tt)
	//var ixField reflect.StructField
	for i := 0; i < 3; i++ {
		ixField := reflectType.Field(i)
		fmt.Printf("%s\n", ixField.Tag)
	}
}
