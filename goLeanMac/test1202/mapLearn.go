package main

import "fmt"

func main() {
	myMap := make(map[int]int)
	myMap[1] = 2
	_, ok := myMap[12]
	if ok {
		fmt.Println("ok")
	} else {
		fmt.Println("notOk")
	}

}
