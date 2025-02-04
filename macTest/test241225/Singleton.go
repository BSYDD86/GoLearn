package main

import (
	"sync"
)

type SingleTon3 struct {
	name string
}

var instance *SingleTon3

var once sync.Once

func GetSingleTon3() *SingleTon3 {
	once.Do(func() {
		instance = &SingleTon3{name: "3"}
	})
	return instance

}

//func main() {
//	s := GetSingleTon3()
//	fmt.Println(s.name)
//}
