package main

import "fmt"

func main() {
	str := "123ab"
	arr := []byte{}

	for _, n := range str {
		arr = append(arr, byte(n))
	}
	fmt.Println(arr)
	fmt.Println(string(arr))
}
