package main

import "fmt"

func main() {
	arr := make([]int, 3)
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[0] = 100
	arr1 := make([]int, 2)
	copy(arr1, arr)
	fmt.Println(arr1)

}
