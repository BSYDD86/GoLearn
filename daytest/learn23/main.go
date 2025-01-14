package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) call() {
	fmt.Println("hello!")
	fmt.Println("this is:", this)
}

func main() {
	user := User{1, "see", 1}
	fmt.Println(user)
	DoFilt(user)
}

func DoFilt(input interface{}) {
	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)
	fmt.Println("----------")
	fmt.Println("type is:", inputType)
	fmt.Println("value is:", inputValue)

	//relfect
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("11111")
		fmt.Printf("%v", field)
		fmt.Printf("%v", value)
	}
}
