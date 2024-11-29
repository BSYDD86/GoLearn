package main

import "fmt"

func main() {
	myMap := map[string]string{
		"asian":  "beijing",
		"europe": "frence",
	}

	fmt.Println(myMap)

	delete(myMap, "asian")

	myMap["europe"] = "helllo"
}
