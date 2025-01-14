package main

import "sync"

type SingleTon struct {
	SingleTon string
}

var once sync.Once
var instance *SingleTon

func GetSingleTon() *SingleTon {
	once.Do(func() {
		instance = &SingleTon{SingleTon: "aaa"}
	})
	return instance
}
