package main

import "fmt"

func main() {
	myMap := make(map[string]string, 10)
	myMap["str"] = "ss"
	for key := range myMap {
		fmt.Println(key)
	}
}
