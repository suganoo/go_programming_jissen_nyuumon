package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.ValueOf(200))
	var v interface{}
	fmt.Println(reflect.ValueOf(v))
}
