package main

import (
	"fmt"
	"reflect"
)

var (
	str []string
)

func main() {
	a := 'a'
	b := "a"
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))

}
